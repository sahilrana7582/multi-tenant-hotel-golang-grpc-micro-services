package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sahilrana7582/multi-tenant-hotel/pkg/errs"
	"github.com/sahilrana7582/multi-tenant-hotel/user-service/models"
	"github.com/sahilrana7582/multi-tenant-hotel/user-service/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	var user models.NewUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return errs.New("Invalid input", err.Error(), http.StatusBadRequest)
	}
	createdUser, err := h.service.CreateUser(r.Context(), &user)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(createdUser); err != nil {
		return errs.ErrInternalServer
	}

	return nil
}
