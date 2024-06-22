package router

import (
	"github.com/douglastenfen/go-rest-api/internal/handlers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/events", handlers.GetEvents).Methods("GET")
	r.HandleFunc("/events/{eventId}", handlers.GetEventById).Methods("GET")
	r.HandleFunc("/events/{eventId}/spots", handlers.GetSpotsByEventId).Methods("GET")
	r.HandleFunc("/events/{eventId}/reserve", handlers.ReserveSpot).Methods("POST")
	r.HandleFunc("/spots", handlers.GetSpots).Methods("GET")

	return r
}
