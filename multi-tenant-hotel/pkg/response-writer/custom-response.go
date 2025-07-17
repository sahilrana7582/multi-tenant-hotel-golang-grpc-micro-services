package responsewriter

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func WriteSuccess(w http.ResponseWriter, statusCode int, message string, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := SuccessResponse{
		Code:    http.StatusText(statusCode),
		Message: message,
		Data:    data,
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return err
	}

	return nil
}
