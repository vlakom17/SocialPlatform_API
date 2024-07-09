package routes

import (
	handlers "api/internal/services"

	"github.com/gorilla/mux"
)

func InitExternalRoutes(router *mux.Router) {
	// Маршрут для работы с external api
	router.HandleFunc("/v1/events/{event_id}", handlers.GetEventHandler).Methods("POST")
}
