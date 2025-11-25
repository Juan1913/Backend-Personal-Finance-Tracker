package errors

import (
	"encoding/json"
	"net/http"
	"time"
)

type ApiError struct {
	Code           string   `json:"code"`
	Status         string   `json:"status"`
	Message        string   `json:"message"`
	DetailMessages []string `json:"detailMessages"`
	TimeStamp      string   `json:"timeStamp"`
	HttpStatus     int      `json:"-"`
}

func NewApiError(code, status, message string, details []string, httpStatus int) error {
	return &ApiError{
		Code:           code,
		Status:         status,
		Message:        message,
		DetailMessages: details,
		TimeStamp:      time.Now().Format(time.RFC3339Nano),
		HttpStatus:     httpStatus,
	}
}

func (e *ApiError) Error() string {
	return e.Message
}

func WriteError(w http.ResponseWriter, apiErr *ApiError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(apiErr.HttpStatus)
	_ = json.NewEncoder(w).Encode(apiErr)
}
