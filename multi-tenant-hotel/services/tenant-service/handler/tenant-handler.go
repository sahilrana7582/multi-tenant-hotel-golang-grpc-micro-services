package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/sahilrana7582/hotel-mgmt/services/tenant-service/models"
	"github.com/sahilrana7582/hotel-mgmt/services/tenant-service/service"
	"github.com/sahilrana7582/multi-tenant-hotel/pkg/errs"
	responsewriter "github.com/sahilrana7582/multi-tenant-hotel/pkg/response-writer"
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

func (h *TenantHandler) GetTenant(w http.ResponseWriter, r *http.Request) error {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	id := chi.URLParam(r, "id")
	tenant, err := h.service.GetTenant(ctx, id)
	if err != nil {
		return err
	}

	return responsewriter.WriteSuccess(w, http.StatusOK, "Tenant retrieved successfully", tenant)
}

func (h *TenantHandler) ListTenants(w http.ResponseWriter, r *http.Request) error {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	tenants, err := h.service.ListTenants(ctx)
	if err != nil {
		return err
	}

	return responsewriter.WriteSuccess(w, http.StatusOK, "Tenants retrieved successfully", tenants)
}

func (h *TenantHandler) UpdateTenant(w http.ResponseWriter, r *http.Request) error {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	id := chi.URLParam(r, "id")
	var input models.UpdateTenantInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return errs.ErrInvalidInput
	}
	tenant, err := h.service.UpdateTenant(ctx, id, &input)
	if err != nil {
		return err
	}

	return responsewriter.WriteSuccess(w, http.StatusOK, "Tenant updated successfully", tenant)
}

func (h *TenantHandler) DeleteTenant(w http.ResponseWriter, r *http.Request) error {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	id := chi.URLParam(r, "id")
	err := h.service.DeleteTenant(ctx, id)
	if err != nil {
		return err
	}

	return responsewriter.WriteSuccess(w, http.StatusOK, "Tenant deleted successfully", nil)
}
