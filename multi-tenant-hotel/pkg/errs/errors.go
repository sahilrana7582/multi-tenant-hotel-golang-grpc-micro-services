package errs

import "net/http"

type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) HTTPStatus() int {
	return e.Status
}

func New(code, message string, httpStatus int) *AppError {
	return &AppError{Code: code, Message: message, Status: httpStatus}
}

func Wrap(err error, code string, status int) *AppError {
	return &AppError{
		Message: err.Error(),
		Code:    code,
		Status:  status,
	}
}

var (

	// General Errors
	ErrBadRequest     = New("bad request", "BAD_REQUEST", http.StatusBadRequest)
	ErrUnauthorized   = New("unauthorized", "UNAUTHORIZED", http.StatusUnauthorized)
	ErrForbidden      = New("forbidden", "FORBIDDEN", http.StatusForbidden)
	ErrNotFound       = New("resource not found", "NOT_FOUND", http.StatusNotFound)
	ErrInternalServer = New("internal server error", "INTERNAL_ERROR", http.StatusInternalServerError)

	// Tenent Service
	ErrDuplicateTenant = New("TENANT_DUPLICATE", "tenant already exists", http.StatusConflict)
	ErrInvalidInput    = New("INVALID_INPUT", "invalid request input", http.StatusBadRequest)
	ErrInternal        = New("INTERNAL_ERROR", "internal server error", http.StatusInternalServerError)
)
