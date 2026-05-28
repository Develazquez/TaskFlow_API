package core

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if payload != nil {
		json.NewEncoder(w).Encode(payload)
	}
}

func RespondError(w http.ResponseWriter, status int, message string) {
	RespondJSON(w, status, ErrorResponse{Message: message})
}

// MapErrorToHTTPStatus ayuda a mapear los errores de dominio a códigos HTTP
func MapErrorToHTTPStatus(err error) int {
	switch err {
	case ErrUserNotFound, ErrProjectNotFound, ErrTaskNotFound:
		return http.StatusNotFound
	case ErrUserAlreadyExists, ErrInvalidInput:
		return http.StatusBadRequest
	case ErrInvalidCredentials, ErrUnauthorized:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
