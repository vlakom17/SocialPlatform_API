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
	router.HandleFunc("/v1/auth/register", handlers.RegisterUserHandler).Methods("POST")
	router.HandleFunc("/v1/auth/edit-user", handlers.EditUserHandler).Methods("PUT")
	router.HandleFunc("/v1/auth/get-user", handlers.GetUserHandler).Methods("GET")

	// Маршруты для работы с мероприятиями
	router.HandleFunc("/v1/events/create", handlers.CreateEventHandler).Methods("POST")
	router.HandleFunc("/v1/events/edit", handlers.EditEventHandler).Methods("PUT")
	router.HandleFunc("/v1/events/get", handlers.GetEventHandler).Methods("GET")

	// Маршруты для работы с аудиторией
	router.HandleFunc("/v1/audience/create", handlers.CreateAudienceHandler).Methods("POST")
	router.HandleFunc("/v1/audience/edit", handlers.EditAudienceHandler).Methods("PUT")
	router.HandleFunc("/v1/audience/get", handlers.GetAudienceHandler).Methods("GET")
	router.HandleFunc("/v1/audience/delete", handlers.DeleteAudienceHandler).Methods("DELETE")
}
