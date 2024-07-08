package handlers

import (
	"api/internal/models"
	"encoding/json"
	"io"
	"net/http"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Пример отправки запроса на сервис авторизации для регистрации пользователя
	resp, err := http.Post(accountServiceURL+"/register", "application/json", r.Body)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}
func EditUserHandler(w http.ResponseWriter, r *http.Request) {
	var req models.EditUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Пример отправки запроса на сервис авторизации для редактирования данных пользователя
	resp, err := http.Post(accountServiceURL+"/edit-user", "application/json", r.Body)
	if err != nil {
		http.Error(w, "Failed to edit user data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var req models.GetUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Пример отправки запроса на сервис авторизации для получения данных о пользователе
	resp, err := http.Post(accountServiceURL+"/get-user", "application/json", r.Body)
	if err != nil {
		http.Error(w, "Failed to get user data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}
