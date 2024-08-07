package main

import (
	"net/http"
	"poll-api/controllers"
	"poll-api/database"
	"poll-api/handler.go"

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

	pollHandler := handler.NewPollHandler(database.DB)
	pollController := controllers.NewPollsController(pollHandler)

	r.Get("/polls", pollController.GetPolls)
	r.Get("/polls/{pollsId}", pollController.GetPollsById)
	r.Post("/polls", pollController.CreatePolls)
	r.Put("/polls/{pollsId}", pollController.UpdatePolls)
	r.Put("/polls/{pollsId}/{option}", pollController.UpdatePollsVote)
	r.Delete("/polls/{pollsId}", pollController.DeletePolls)

	http.ListenAndServe(":4343", r)
}
