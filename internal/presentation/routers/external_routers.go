package routes

import (
	handlers "api/internal/services"

	"github.com/gorilla/mux"
)

func InitExternalRoutes(router *mux.Router) {
	// Маршрут для работы с external api
	router.HandleFunc("/external_api/{endpoint}", handlers.ProxyExternalAPI).Methods("GET", "POST", "PUT", "DELETE")
}
