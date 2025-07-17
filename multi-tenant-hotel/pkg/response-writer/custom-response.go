package responsewriter

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sahilrana7582/multi-tenant-hotel/pkg/errs"
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

func WriteError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError

	var sc interface{ HTTPStatus() int }
	if errors.As(err, &sc) {
		status = sc.HTTPStatus()
	}

	var resp *errs.AppError
	if appErr, ok := err.(*errs.AppError); ok {
		resp = appErr
	} else {
		resp = errs.Wrap(err, errs.ErrInternalServer.Code, errs.ErrInternalServer.Status)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    resp.Code,
		"message": resp.Message,
	})
}
