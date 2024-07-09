package models

// SessionResponse структура для ответа от сервиса авторизации
type SessionResponse struct {
	Valid bool `json:"valid"`
}

// CreateSessionRequest структура для создания сессии
type CreateSessionRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// SessionToken структура для запроса завершения сессии
type SessionToken struct {
	Token string `json:"token"`
}

// RequestPayload представляет структуру данных, ожидаемую от внешнего API
type RequestPayload struct {
	Data string `json:"data"`
}

// ResponsePayload представляет структуру данных, которую мы отправим в ответе
type ResponsePayload struct {
	Message string `json:"message"`
}
