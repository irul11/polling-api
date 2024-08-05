package main

import (
	"net/http"
	"poll-api/controllers"
	"poll-api/database"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func init() {
	database.LoadDatabase()
}

func main() {
	defer database.DB.Close()

	r := chi.NewRouter()
	corsOptions := cors.Options{
		AllowedMethods: []string{"GET", "PUT", "DELETE", "POST", "OPTIONS"},
	}

	r.Use(cors.Handler(corsOptions))
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is API endpoint for poll app"))
	})
	r.Get("/polls", controllers.GetPolls)
	r.Get("/polls/{pollsId}", controllers.GetPollsById)
	r.Post("/polls", controllers.CreatePolls)
	r.Put("/polls/{pollsId}", controllers.UpdatePolls)
	r.Put("/polls/{pollsId}/{option}", controllers.UpdatePollsVote)
	r.Delete("/polls/{pollsId}", controllers.DeletePolls)

	http.ListenAndServe(":4343", r)
}
