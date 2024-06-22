package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/douglastenfen/go-rest-api/internal/repository"
	"github.com/gorilla/mux"
)

func GetSpots(w http.ResponseWriter, r *http.Request) {
	data, err := repository.LoadData()
	if err != nil {
		http.Error(w, "Unable to load data", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data.Spots)
}

func GetSpotsByEventId(w http.ResponseWriter, r *http.Request) {
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

	var eventSpots []string
	for _, spot := range data.Spots {
		if spot.EventID == eventId {
			eventSpots = append(eventSpots, spot.Name)
		}
	}

	if len(eventSpots) == 0 {
		http.Error(w, "No spots found for this event", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(eventSpots)
}

func ReserveSpot(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	eventId, err := strconv.Atoi(params["eventId"])
	if err != nil {
		log.Println("Invalid event ID")
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	var reserveRequest struct {
		SpotName string `json:"spot_name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reserveRequest); err != nil {
		log.Println("Invalid request payload")
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	data, err := repository.LoadData()
	if err != nil {
		log.Println("Unable to load data")
		http.Error(w, "Unable to load data", http.StatusInternalServerError)
		return
	}

	for i, spot := range data.Spots {
		if spot.EventID == eventId && spot.Name == reserveRequest.SpotName {
			if spot.Status == "reserved" {
				log.Println("Spot already reserved")
				http.Error(w, "Spot already reserved", http.StatusBadRequest)
				return
			}
			data.Spots[i].Status = "reserved"
			json.NewEncoder(w).Encode(spot)
			log.Println("Spot reserved successfully")
			return
		}
	}

	log.Println("Spot not found")
	http.Error(w, "Spot not found", http.StatusNotFound)
}
