package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/douglastenfen/go-rest-api/internal/repository"
	"github.com/gorilla/mux"
)

func GetEvents(w http.ResponseWriter, r *http.Request) {
	data, err := repository.LoadData()
	if err != nil {
		http.Error(w, "Unable to load data", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data.Events)
}

func GetEventById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	eventId, err := strconv.Atoi(params["eventId"])
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	data, err := repository.LoadData()
	if err != nil {
		http.Error(w, "Unable to load data", http.StatusInternalServerError)
		return
	}

	for _, event := range data.Events {
		if event.ID == eventId {
			json.NewEncoder(w).Encode(event)
			return
		}
	}

	http.Error(w, "Event not found", http.StatusNotFound)
}
