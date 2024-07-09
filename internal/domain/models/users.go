package models

import "github.com/google/uuid"

// UserUpdateRequest представляет структуру запроса для редактирования данных пользователя
type UserUpdateRequest struct {
}

// UserRegistrationRequest представляет структуру запроса для регистрации пользователя
type UserRegistrationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserResponse представляет структуру ответа с данными о пользователе от аккаунт-сервиса
// UserResponse представляет структуру ответа от аккаунт-сервиса
type UserResponse struct {
	UserId   uuid.UUID `json:"user_id"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
	Roles    []string  `json:"roles"`
	Status   string    `json:"status"`
	Message  string    `json:"message"`
}
