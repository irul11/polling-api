package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"poll-api/handler.go"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Polls struct {
	Name string `json:"name"`
}

func GetPolls(w http.ResponseWriter, r *http.Request) {
	polls, err := handler.GetPolls()
	if err != nil {
		http.Error(w, "Error on querying polls data", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(polls)
	if err != nil {
		http.Error(w, "Error marshalling polls data", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(response)
	if err != nil {
		log.Printf("Error writing response: %v", err)
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}

func UpdatePollsVote(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	option := queryParams.Get("option")

	pollsIdParam := chi.URLParam(r, "pollsId")
	pollsId, err := strconv.Atoi(pollsIdParam)

	if err != nil {
		// Handle invalid ID error
		http.Error(w, "Invalid polls ID", http.StatusBadRequest)
		return
	}

	err = handler.UpdatePollsVote(pollsId, option)
	if err != nil {
		http.Error(w, "Error updating vote", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Updating vote succesful"))
}
