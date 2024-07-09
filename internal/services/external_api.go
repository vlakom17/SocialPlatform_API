package handlers

import (
	"encoding/json"
	"net/http"
)

const (
	eventServiceURL   = "http://localhost:8080" // URL сервиса управления мероприятиями
	accountServiceURL = "http://localhost:8080" //URL сервиса аккаунта
)

func HandleAPIRequest(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что метод запроса POST
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Декодируем JSON из тела запроса
	var requestData struct {
		Service_Method string `json:"service_method"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Ошибка при чтении JSON", http.StatusBadRequest)
		return
	}

	// Определяем, какую ручку активировать на основе данных из запроса
	switch requestData.Service_Method {
	case "events_post":
		// Вызываем функцию для обработки конкретной ручки
		CreateEventHandler(w, r)
	case "events_put":
		// Вызываем функцию для обработки другой ручки
		UpdateEventHandler(w, r)
	case "events_get":
		GetEventHandler(w, r)
	case "rooms_post":
		CreateRoomHandler(w, r)
	case "rooms_get":
		GetRoomHandler(w, r)
	case "rooms_delete":
		DeleteRoomHandler(w, r)
	case "users_post":
		RegisterUserHandler(w, r)
	case "users_put":
		EditUserHandler(w, r)
	case "users_get":
		GetUserDataHandler(w, r)
	default:
		http.Error(w, "Unknown Method", http.StatusNotFound)
		return
	}
}

/*
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
*/
