package handlers

import (
	"api/internal/domain/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// CreateEvent отправляет POST-запрос на создание мероприятия к event-сервису
func CreateEvent(event models.Event, eventServiceURL string) (*models.EventResponse, error) {

	// Преобразуем структуру Event в JSON
	eventJSON, err := json.Marshal(event)
	if err != nil {
		return nil, fmt.Errorf("error marshalling event: %v", err)
	}

	// Создаем HTTP-клиент
	client := &http.Client{}

	// Создаем новый POST-запрос
	req, err := http.NewRequest("POST", eventServiceURL+"/v1/events", bytes.NewBuffer(eventJSON))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Устанавливаем заголовки
	req.Header.Set("Content-Type", "application/json")

	// Отправляем запрос
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем код ответа
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("received non-200 response: %v - %s", resp.Status, string(body))
	}

	// Декодируем ответ
	var eventResp models.EventResponse
	if err := json.NewDecoder(resp.Body).Decode(&eventResp); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}

	return &eventResp, nil
}

// CreateEventHandler обрабатывает запрос на создание нового мероприятия
func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	// Декодируем JSON из тела запроса в структуру Event
	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Вызываем функцию для создания мероприятия
	eventResp, err := CreateEvent(event, eventServiceURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем ответ клиенту в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(eventResp)
}

// UpdateEvent отправляет PUT-запрос на редактирование мероприятия к event-сервису по его ID типа uuid.UUID
func UpdateEvent(eventID uuid.UUID, event models.Event, eventServiceURL string) (*models.EditEventResponse, error) {
	// Преобразуем структуру Event в JSON
	eventJSON, err := json.Marshal(event)
	if err != nil {
		return nil, fmt.Errorf("error marshalling event: %v", err)
	}

	// Создаем HTTP-клиент
	client := &http.Client{}

	// Формируем полный URL с учетом eventID
	fullURL := fmt.Sprintf("%s/v1/events/%s", eventServiceURL, eventID.String())

	// Создаем новый PUT-запрос
	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(eventJSON))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Устанавливаем заголовки
	req.Header.Set("Content-Type", "application/json")

	// Отправляем запрос
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем код ответа
	if resp.StatusCode != http.StatusOK {
		// Читаем тело ответа для более подробного сообщения об ошибке
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("received non-200 response: %v - %s", resp.Status, string(body))
	}

	// Декодируем ответ
	var eventResp models.EditEventResponse
	if err := json.NewDecoder(resp.Body).Decode(&eventResp); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}

	return &eventResp, nil
}

// UpdateEventHandler обрабатывает запрос на редактирование мероприятия
func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем ID мероприятия из пути
	vars := mux.Vars(r)
	eventID, err := uuid.Parse(vars["event_id"])
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	// Декодируем JSON из тела запроса в структуру Event
	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Вызываем функцию для обновления мероприятия
	eventResp, err := UpdateEvent(eventID, event, eventServiceURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем ответ клиенту в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(eventResp)
}

// GetEventByID отправляет GET-запрос на получение данных о мероприятии по его ID типа uuid.UUID
func GetEventByID(eventID uuid.UUID, eventServiceURL string) (*models.EditEventResponse, error) {
	// Формируем полный URL с учетом eventID
	fullURL := fmt.Sprintf("%s/v1/events/%s", eventServiceURL, eventID.String())

	// Создаем HTTP-клиент
	client := &http.Client{}

	// Создаем новый GET-запрос
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Устанавливаем заголовки
	req.Header.Set("Content-Type", "application/json")

	// Отправляем запрос
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем код ответа
	if resp.StatusCode != http.StatusOK {
		// Читаем тело ответа для более подробного сообщения об ошибке
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("received non-200 response: %v - %s", resp.Status, string(body))
	}

	// Декодируем ответ
	var eventResp models.EditEventResponse
	if err := json.NewDecoder(resp.Body).Decode(&eventResp); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}

	return &eventResp, nil
}

// GetEventHandler обрабатывает запрос на получение данных о мероприятии
func GetEventHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем ID мероприятия из пути
	vars := mux.Vars(r)
	eventID, err := uuid.Parse(vars["event_id"])
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	// Вызываем функцию для получения данных о мероприятии
	eventResp, err := GetEventByID(eventID, eventServiceURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем ответ клиенту в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(eventResp)
}

// CreateRoom отправляет POST-запрос на комнаты к event-сервису
func CreateRoom(room models.Room, eventServiceURL string) (*models.RoomResponse, error) {

	// Преобразуем структуру Event в JSON
	roomJSON, err := json.Marshal(room)
	if err != nil {
		return nil, fmt.Errorf("error marshalling room: %v", err)
	}

	// Создаем HTTP-клиент
	client := &http.Client{}

	// Создаем новый POST-запрос
	req, err := http.NewRequest("POST", eventServiceURL+"/v1/rooms", bytes.NewBuffer(roomJSON))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Устанавливаем заголовки
	req.Header.Set("Content-Type", "application/json")

	// Отправляем запрос
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем код ответа
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("received non-200 response: %v - %s", resp.Status, string(body))
	}

	// Декодируем ответ
	var roomResp models.RoomResponse
	if err := json.NewDecoder(resp.Body).Decode(&roomResp); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}

	return &roomResp, nil
}

// CreateRoomHandler обрабатывает запрос на создание комнаты
func CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	// Инициализируем переменную для хранения данных комнаты
	var room models.Room

	// Декодируем JSON из тела запроса в структуру Room
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Вызываем функцию для создания комнаты
	roomResp, err := CreateRoom(room, eventServiceURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем ответ клиенту в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roomResp)
}

// GetRoomByID отправляет GET-запрос на получение данных о комнате по её ID типа uuid.UUID
func GetRoomByID(roomID uuid.UUID, eventServiceURL string) (*models.EditRoomResponse, error) {
	// Формируем полный URL с учетом eventID
	fullURL := fmt.Sprintf("%s/v1/rooms/%s", eventServiceURL, roomID.String())

	// Создаем HTTP-клиент
	client := &http.Client{}

	// Создаем новый GET-запрос
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Устанавливаем заголовки
	req.Header.Set("Content-Type", "application/json")

	// Отправляем запрос
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем код ответа
	if resp.StatusCode != http.StatusOK {
		// Читаем тело ответа для более подробного сообщения об ошибке
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("received non-200 response: %v - %s", resp.Status, string(body))
	}

	// Декодируем ответ
	var roomResp models.EditRoomResponse
	if err := json.NewDecoder(resp.Body).Decode(&roomResp); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}

	return &roomResp, nil
}

// GetRoomByIDHandler обрабатывает запрос на получение данных о комнате по её ID
func GetRoomHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем параметр из пути (ID комнаты)
	vars := mux.Vars(r)
	roomID := vars["room_id"]

	// Преобразуем roomID в тип uuid.UUID
	uuidRoomID, err := uuid.Parse(roomID)
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}

	// Вызываем функцию для получения данных о комнате
	roomResp, err := GetRoomByID(uuidRoomID, eventServiceURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем ответ клиенту в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roomResp)
}

// DeleteRoom отправляет DELETE-запрос на удаление комнаты по её ID к event-сервису
func DeleteRoom(roomID uuid.UUID, eventServiceURL string) error {
	// Формируем полный URL с учетом roomID
	fullURL := fmt.Sprintf("%s/v1/rooms/%s", eventServiceURL, roomID.String())

	// Создаем HTTP-клиент
	client := &http.Client{}

	// Создаем новый DELETE-запрос
	req, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	// Отправляем запрос
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем код ответа
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 response: %v", resp.Status)
	}

	return nil
}

// DeleteRoomHandler обрабатывает запрос на удаление комнаты по её ID
func DeleteRoomHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем параметр из пути (ID комнаты)
	vars := mux.Vars(r)
	roomID := vars["room_id"]

	// Преобразуем roomID в тип uuid.UUID
	uuidRoomID, err := uuid.Parse(roomID)
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}

	// Вызываем функцию для удаления комнаты
	err = DeleteRoom(uuidRoomID, eventServiceURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный статус
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Room %s deleted successfully", roomID)
}
