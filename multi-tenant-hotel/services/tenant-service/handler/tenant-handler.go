package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/sahilrana7582/hotel-mgmt/services/tenant-service/models"
	"github.com/sahilrana7582/hotel-mgmt/services/tenant-service/service"
	"github.com/sahilrana7582/multi-tenant-hotel/pkg/errs"
)

type TenantHandler struct {
	service service.TenantService
}

func NewTenantHandler(r service.TenantService) *TenantHandler {
	return &TenantHandler{service: r}
}

func (h *TenantHandler) CreateTenant(w http.ResponseWriter, r *http.Request) error {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var input models.CreateTenantInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return errs.ErrInvalidInput
	}

	tenant, err := h.service.RegisterTenant(ctx, &input)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(tenant); err != nil {
		return errs.ErrInternalServer
	}

	return nil
}
