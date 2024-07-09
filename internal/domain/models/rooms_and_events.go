package models

import (
	"time"

	"github.com/google/uuid"
)

// CreateEventRequest структура для запроса создания мероприятия
type Event struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CreatorId   uuid.UUID `json:"creator_id"`
	RoomId      uuid.UUID `json:"room_id"`
}

// EventResponse представляет структуру ответа от event-сервиса
type EventResponse struct {
	ID      uuid.UUID `json:"event_id"`
	Message string    `json:"message,omitempty"`
}
type EditEventResponse struct {
	ID          uuid.UUID `json:"event_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CreatorId   uuid.UUID `json:"creator_id"`
	RoomId      uuid.UUID `json:"room_id"`
}

type Room struct {
	Title string `json:"title"`
}

type RoomResponse struct {
	RoomId  uuid.UUID `json:"room_id"`
	Message string    `json:"message,omitempty"`
}

type EditRoomResponse struct {
	RoomID uuid.UUID `json:"room_id"`
	Title  string    `json:"title"`
}
