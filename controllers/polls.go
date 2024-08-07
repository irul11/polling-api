package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"poll-api/handler.go"
	"poll-api/models"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PollsController struct {
	Handler handler.PollHandler
}

func NewPollsController(ph handler.PollHandler) *PollsController {
	return &PollsController{
		Handler: ph,
	}
}

func (pc *PollsController) GetPolls(w http.ResponseWriter, r *http.Request) {
	polls, err := pc.Handler.GetPolls()
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

func (pc *PollsController) GetPollsById(w http.ResponseWriter, r *http.Request) {
	pollsIdParam := chi.URLParam(r, "pollsId")
	pollsId, err := strconv.Atoi(pollsIdParam)
	if err != nil {
		// Handle invalid ID error
		http.Error(w, "Invalid polls ID", http.StatusBadRequest)
		return
	}

	polls, err := pc.Handler.GetPollsById(pollsId)
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

func (pc *PollsController) CreatePolls(w http.ResponseWriter, r *http.Request) {
	var polls models.Polls

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&polls)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = pc.Handler.CreatePolls(polls)
	if err != nil {
		// Handle error
		http.Error(w, "Failed insert data", http.StatusInternalServerError)
		return
	}
}

func (pc *PollsController) UpdatePolls(w http.ResponseWriter, r *http.Request) {
	var polls models.Polls
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&polls)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pollsIdParam := chi.URLParam(r, "pollsId")
	pollsId, err := strconv.Atoi(pollsIdParam)
	if err != nil {
		// Handle invalid ID error
		http.Error(w, "Invalid polls ID", http.StatusBadRequest)
		return
	}

	err = pc.Handler.UpdatePolls(pollsId, polls)
	if err != nil {
		http.Error(w, "Error updating polls", http.StatusInternalServerError)
		return
	}
}

func (pc *PollsController) UpdatePollsVote(w http.ResponseWriter, r *http.Request) {
	option := chi.URLParam(r, "option")
	if option != "a" && option != "b" {
		http.Error(w, "Invalid request option", http.StatusBadRequest)
		return
	}

	pollsIdParam := chi.URLParam(r, "pollsId")
	pollsId, err := strconv.Atoi(pollsIdParam)

	if err != nil {
		// Handle invalid ID error
		http.Error(w, "Invalid polls ID", http.StatusBadRequest)
		return
	}

	err = pc.Handler.UpdatePollsVote(pollsId, option)
	if err != nil {
		http.Error(w, "Error updating vote", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Updating vote succesful"))
}

func (pc *PollsController) DeletePolls(w http.ResponseWriter, r *http.Request) {
	pollsIdParam := chi.URLParam(r, "pollsId")
	pollsId, err := strconv.Atoi(pollsIdParam)

	if err != nil {
		// Handle invalid ID error
		http.Error(w, "Invalid polls ID", http.StatusBadRequest)
		return
	}

	err = pc.Handler.DeletePolls(pollsId)
	if err != nil {
		http.Error(w, "Error deleting polls", http.StatusInternalServerError)
	}
}
