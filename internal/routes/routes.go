package routes

import (
	"api/internal/handlers"

	"github.com/gorilla/mux"
)

// InitRoutes инициализирует маршруты
func InitRoutes(router *mux.Router) {
	// Маршруты для создания и завершения сессии авторизации
	router.HandleFunc("/v1/auth/create-session", handlers.CreateSessionHandler).Methods("POST")
	router.HandleFunc("/v1/auth/end-session", handlers.EndSessionHandler).Methods("POST")

	// Маршруты для регистрации пользователя, редактирования данных и получения данных о пользователе
	router.HandleFunc("/v1/users", handlers.RegisterUserHandler).Methods("POST")
	router.HandleFunc("/v1/users/{user_id}", handlers.EditUserHandler).Methods("PUT")
	router.HandleFunc("/v1/users/{user_id}", handlers.GetUserDataHandler).Methods("GET")

	// Маршруты для работы с мероприятиями
	router.HandleFunc("/v1/events", handlers.CreateEventHandler).Methods("POST")
	router.HandleFunc("/v1/events/{event_id}", handlers.UpdateEventHandler).Methods("PUT")
	router.HandleFunc("/v1/events/{event_id}", handlers.GetEventHandler).Methods("GET")

	// Маршруты для работы с аудиторией
	router.HandleFunc("/v1/rooms", handlers.CreateRoomHandler).Methods("POST")
	router.HandleFunc("/v1/rooms/{room_id}", handlers.GetRoomHandler).Methods("GET")
	router.HandleFunc("/v1/rooms/{room_id}", handlers.DeleteRoomHandler).Methods("DELETE")
}
