package routes

import (
	handlers "api/internal/services"

	"github.com/gorilla/mux"
)

func InitRoomRoutes(router *mux.Router) {

	// Маршруты для работы с аудиторией
	router.HandleFunc("/v1/rooms", handlers.CreateRoomHandler).Methods("POST")
	router.HandleFunc("/v1/rooms/{room_id}", handlers.GetRoomHandler).Methods("GET")
	router.HandleFunc("/v1/rooms/{room_id}", handlers.DeleteRoomHandler).Methods("DELETE")
}
