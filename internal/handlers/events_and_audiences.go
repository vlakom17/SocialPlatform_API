package handlers

import (
	"api/internal/models"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// Маршрут для создания мероприятия
func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	var req models.CreateEventRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		http.Error(w, "Failed to encode request", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post(eventServiceURL+"/create-event", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		http.Error(w, "Failed to create event", http.StatusInternalServerError)
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

// Маршрут для редактирования мероприятия
func EditEventHandler(w http.ResponseWriter, r *http.Request) {
	var req models.EditEventRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		http.Error(w, "Failed to encode request", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post(eventServiceURL+"/edit-event", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		http.Error(w, "Failed to edit event", http.StatusInternalServerError)
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

// Маршрут для запроса данных о мероприятии
func GetEventHandler(w http.ResponseWriter, r *http.Request) {
	var req models.GetEventRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		http.Error(w, "Failed to encode request", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post(eventServiceURL+"/get-event", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		http.Error(w, "Failed to get event data", http.StatusInternalServerError)
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

// Маршрут для создания аудитории
func CreateAudienceHandler(w http.ResponseWriter, r *http.Request) {
	var req models.CreateAudienceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		http.Error(w, "Failed to encode request", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post(eventServiceURL+"/create-audience", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		http.Error(w, "Failed to create audience", http.StatusInternalServerError)
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

// Маршрут для редактирования аудитории
func EditAudienceHandler(w http.ResponseWriter, r *http.Request) {
	var req models.EditAudienceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		http.Error(w, "Failed to encode request", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post(eventServiceURL+"/edit-audience", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		http.Error(w, "Failed to edit audience", http.StatusInternalServerError)
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

// Маршрут для запроса данных об аудитории
func GetAudienceHandler(w http.ResponseWriter, r *http.Request) {
	var req models.GetAudienceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		http.Error(w, "Failed to encode request", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post(eventServiceURL+"/get-audience", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		http.Error(w, "Failed to get audience data", http.StatusInternalServerError)
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

// Маршрут для удаления аудитории
func DeleteAudienceHandler(w http.ResponseWriter, r *http.Request) {
	var req models.DeleteAudienceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		http.Error(w, "Failed to encode request", http.StatusInternalServerError)
		return
	}

	httpReq, err := http.NewRequest("DELETE", eventServiceURL+"/delete-audience", bytes.NewBuffer(reqBody))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		http.Error(w, "Failed to delete audience", http.StatusInternalServerError)
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
