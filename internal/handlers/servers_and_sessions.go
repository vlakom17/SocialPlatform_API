package handlers

import (
	"api/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	authServiceURL    = "http://localhost:4000" // URL сервиса авторизации
	eventServiceURL   = "http://localhost:8080" // URL сервиса управления мероприятиями
	accountServiceURL = "http://localhost:6000" //URL сервиса аккаунта
)

// Маршрут для создания сессии авторизации
func CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
	var req models.CreateSessionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		http.Error(w, "Failed to encode request body", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post(authServiceURL+"/create-session", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
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

// Маршрут для завершения сессии авторизации
func EndSessionHandler(w http.ResponseWriter, r *http.Request) {
	var req models.SessionToken
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		http.Error(w, "Failed to encode request body", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post(authServiceURL+"/end-session", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		http.Error(w, "Failed to end session", http.StatusInternalServerError)
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

func authHandler(w http.ResponseWriter, r *http.Request) {
	// Пример получения user cookie
	cookie, err := r.Cookie("user_cookie_name")
	if err != nil {
		http.Error(w, "Failed to get user cookie", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received user cookie: %s=%s\n", cookie.Name, cookie.Value)

	// Пример получения auth token из заголовка Authorization
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Auth token is missing", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received auth token: %s\n", authHeader)

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Received user cookie and auth token successfully"))
}
func externalAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка метода запроса (только POST)
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Декодирование JSON данных из тела запроса
	var reqPayload models.RequestPayload
	if err := json.NewDecoder(r.Body).Decode(&reqPayload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Логирование полученных данных
	fmt.Printf("Received data: %s\n", reqPayload.Data)

	// Обработка данных и создание ответа
	response := models.ResponsePayload{
		Message: "Data received successfully: " + reqPayload.Data,
	}

	// Кодирование ответа в JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
