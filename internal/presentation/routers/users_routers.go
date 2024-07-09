package routes

import (
	handlers "api/internal/services"

	"github.com/gorilla/mux"
)

func InitUserRoutes(router *mux.Router) {

	// Маршруты для регистрации пользователя, редактирования данных и получения данных о пользователе
	router.HandleFunc("/v1/users", handlers.RegisterUserHandler).Methods("POST")
	router.HandleFunc("/v1/users/{user_id}", handlers.EditUserHandler).Methods("PUT")
	router.HandleFunc("/v1/users/{user_id}", handlers.GetUserDataHandler).Methods("GET")
}
