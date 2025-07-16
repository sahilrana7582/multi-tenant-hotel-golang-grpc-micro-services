package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sahilrana7582/multi-tenant-hotel/auth-service/models"
	"github.com/sahilrana7582/multi-tenant-hotel/auth-service/service"
	"github.com/sahilrana7582/multi-tenant-hotel/pkg/errs"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Login request received")
	var loginReq models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		return errs.New("Invalid input", err.Error(), http.StatusBadRequest)
	}

	fmt.Println("1")

	token, err := h.service.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		return err
	}

	fmt.Println("2")
	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, cookie)

	loginResp := models.LoginResponse{
		Token:   token,
		Message: "Login successful",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println("3")

	if err := json.NewEncoder(w).Encode(loginResp); err != nil {
		return errs.ErrInternalServer
	}
	return nil
}
