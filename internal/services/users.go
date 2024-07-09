package handlers

import (
	"api/internal/domain/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func RegisterUser(username, password string) (*models.UserRegistrationResponse, error) {
	// Создаем запрос для регистрации пользователя
	reqBody := &models.UserRegistrationRequest{
		Username: username,
		Password: password,
	}

	// Преобразуем запрос в JSON
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %v", err)
	}

	// Делаем HTTP POST запрос к аккаунт-сервису
	resp, err := http.Post(accountServiceURL+"/v1/users", "application/json", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return nil, fmt.Errorf("error making request to account service: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем код состояния HTTP ответа
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %v", resp.StatusCode)
	}

	// Декодируем ответ от аккаунт-сервиса
	var registrationResp models.UserRegistrationResponse
	if err := json.NewDecoder(resp.Body).Decode(&registrationResp); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}

	return &registrationResp, nil
}

// RegisterUserHandler обрабатывает запрос на регистрацию нового пользователя
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	// Декодируем JSON из тела запроса в структуру UserRegistrationRequest
	var reqBody models.UserRegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Вызываем функцию для регистрации пользователя
	registrationResp, err := RegisterUser(reqBody.Username, reqBody.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем ответ клиенту в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(registrationResp)
}

// UpdateUser отправляет PUT-запрос на изменение данных пользователя к аккаунт-сервису
func UpdateUser(userID uuid.UUID, updateUserRequest models.UserUpdateRequest) (*models.UserUpdateResponse, error) {
	// Преобразуем данные для обновления в JSON
	reqBodyBytes, err := json.Marshal(updateUserRequest)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %v", err)
	}

	// Создаем HTTP PUT запрос к аккаунт-сервису
	reqURL := fmt.Sprintf("%s/v1/users/%s", accountServiceURL, userID)
	req, err := http.NewRequest(http.MethodPut, reqURL, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request to account service: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем код состояния HTTP ответа
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %v", resp.StatusCode)
	}

	// Декодируем ответ от аккаунт-сервиса
	var updateResp models.UserUpdateResponse
	if err := json.NewDecoder(resp.Body).Decode(&updateResp); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}

	return &updateResp, nil
}

// EditUserHandler отправляет PUT-запрос на изменение данных пользователя к аккаунт-сервису
func EditUserHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(mux.Vars(r)["user_id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var reqBody models.UserUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := UpdateUser(userID, reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func GetUserData(userID uuid.UUID) (*models.UserResponse, error) {
	// Создаем URL с идентификатором пользователя
	url := fmt.Sprintf(accountServiceURL+"/v1/users/%s", userID)

	// Делаем HTTP GET запрос к аккаунт-сервису
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request to account service: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем код состояния HTTP ответа
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %v", resp.StatusCode)
	}

	// Декодируем ответ от аккаунт-сервиса
	var userResp models.UserResponse
	if err := json.NewDecoder(resp.Body).Decode(&userResp); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}

	return &userResp, nil
}

// GetUserDataHandler обрабатывает запрос на получение данных пользователя
func GetUserDataHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем user_id из пути запроса
	userID, err := uuid.Parse(mux.Vars(r)["user_id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Вызываем функцию для получения данных пользователя
	userResp, err := GetUserData(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем ответ клиенту в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userResp)
}
