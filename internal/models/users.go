package models

import "github.com/google/uuid"

type RegisterUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type EditUserRequest struct {
	User_ID  uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

type GetUserRequest struct {
	UserID uuid.UUID `json:"userId"`
}
