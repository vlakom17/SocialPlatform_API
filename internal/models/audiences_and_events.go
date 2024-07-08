package models

import (
	"time"

	"github.com/google/uuid"
)

// CreateEventRequest структура для запроса создания мероприятия
type CreateEventRequest struct {
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Date         time.Time `json:"date"`
	Organizer_ID uuid.UUID `json:"organizer_id"`
	Audience_ID  uuid.UUID `json:"audience_id"`
	// Дополнительные поля, если необходимо
}

// EditEventRequest структура для запроса редактирования мероприятия
type EditEventRequest struct {
	Event_ID    uuid.UUID `json:"event_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Audience_ID uuid.UUID `json:"audience_id"`
	// Дополнительные поля, которые можно редактировать
}

// GetEventRequest структура для запроса данных о мероприятии
type GetEventRequest struct {
	Event_ID uuid.UUID `json:"event_id"`
}

// CreateAudienceRequest структура для запроса создания аудитории
type CreateAudienceRequest struct {
	Event_ID uuid.UUID `json:"event_id"`
	Audience string    `json:"audience"`
}

// EditAudienceRequest структура для запроса редактирования аудитории
type EditAudienceRequest struct {
	Event_ID    uuid.UUID `json:"event_id"`
	Audience_ID uuid.UUID `json:"audience_id"`
	Audience    string    `json:"audience"`
}

// GetAudienceRequest структура для запроса данных об аудитории
type GetAudienceRequest struct {
	Audience_ID uuid.UUID `json:"audience_id"`
}

// DeleteAudienceRequest структура для запроса удаления аудитории
type DeleteAudienceRequest struct {
	Audience_ID uuid.UUID `json:"audience_id"`
}
