package routes

import (
	handlers "api/internal/services"

	"github.com/gorilla/mux"
)

func InitExternalRoutes(router *mux.Router) {
	// Маршрут для работы с external api
	router.HandleFunc("http://localhost:8080/v1/gateway", handlers.HandleAPIRequest).Methods("POST")
}
