package routes

import (
	handlers "api/internal/services"

	"github.com/gorilla/mux"
)

func InitEventsRoutes(router *mux.Router) {

	// Маршруты для работы с мероприятиями
	router.HandleFunc("/v1/events", handlers.CreateEventHandler).Methods("POST")
	router.HandleFunc("/v1/events/{event_id}", handlers.UpdateEventHandler).Methods("PUT")
	router.HandleFunc("/v1/events/{event_id}", handlers.GetEventHandler).Methods("GET")
}
