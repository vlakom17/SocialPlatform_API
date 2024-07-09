package models

import "github.com/google/uuid"

// UserUpdateRequest представляет структуру запроса для редактирования данных пользователя
type UserUpdateRequest struct {
}

// UserUpdateResponse представляет структуру ответа от аккаунт-сервиса
type UserUpdateResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// UserResponse представляет структуру ответа с данными о пользователе от аккаунт-сервиса
type UserResponse struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
}

// UserRegistrationRequest представляет структуру запроса для регистрации пользователя
type UserRegistrationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserRegistrationResponse представляет структуру ответа от аккаунт-сервиса
type UserRegistrationResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
