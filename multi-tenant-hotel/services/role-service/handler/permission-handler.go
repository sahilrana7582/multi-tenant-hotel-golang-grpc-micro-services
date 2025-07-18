package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sahilrana7582/multi-tenant-hotel/pkg/auth"
	"github.com/sahilrana7582/multi-tenant-hotel/pkg/errs"
	responsewriter "github.com/sahilrana7582/multi-tenant-hotel/pkg/response-writer"
	"github.com/sahilrana7582/multi-tenant-hotel/role-service/models"
	"github.com/sahilrana7582/multi-tenant-hotel/role-service/service"
)

type PermissionHandler struct {
	permissionService service.PermissionService
}

func NewPermissionHandler(permissionService service.PermissionService) *PermissionHandler {
	return &PermissionHandler{
		permissionService: permissionService,
	}
}

func (h *PermissionHandler) GivePermissionToRole(w http.ResponseWriter, r *http.Request) error {

	tenantID := auth.GetTenantID(r)
	if tenantID == "" {
		return errs.New("invalid tenant ID", "invalid tenant ID", http.StatusBadRequest)
	}

	var newPermission models.NewPermission

	if err := json.NewDecoder(r.Body).Decode(&newPermission); err != nil {
		return errs.New("invalid request body", "invalid request body", http.StatusBadRequest)
	}

	permission, err := h.permissionService.GivePermissionToRole(tenantID, &newPermission)
	if err != nil {
		return err
	}

	return responsewriter.WriteSuccess(w, http.StatusCreated, "Permission Granted!", permission)
}

func (h *PermissionHandler) GetPermissionsByRole(w http.ResponseWriter, r *http.Request) error {
	tenantID := auth.GetTenantID(r)
	if tenantID == "" {
		return errs.New("invalid tenant ID", "invalid tenant ID", http.StatusBadRequest)
	}
	roleID := r.URL.Query().Get("role_id")
	if roleID == "" {
		return errs.New("role_id is required", "role_id is required", http.StatusBadRequest)
	}
	permissions, err := h.permissionService.GetPermissionsByRole(tenantID, roleID)
	if err != nil {
		return err
	}
	return responsewriter.WriteSuccess(w, http.StatusOK, "Permissions Fetched Successfully!", permissions)
}

func (h *PermissionHandler) GetAllRolesPermissions(w http.ResponseWriter, r *http.Request) error {
	tenantID := auth.GetTenantID(r)
	if tenantID == "" {
		return errs.New("invalid tenant ID", "invalid tenant ID", http.StatusBadRequest)
	}
	permissions, err := h.permissionService.GetAllRolesPermissions(tenantID)
	if err != nil {
		return err
	}
	return responsewriter.WriteSuccess(w, http.StatusOK, "Permissions Fetched Successfully!", permissions)
}

func (h *PermissionHandler) RemovePermissionFromRole(w http.ResponseWriter, r *http.Request) error {
	tenantID := auth.GetTenantID(r)
	if tenantID == "" {
		return errs.New("invalid tenant ID", "invalid tenant ID", http.StatusBadRequest)
	}
	permissionID := chi.URLParam(r, "id")
	if permissionID == "" {
		return errs.New("permission_id is required", "permission_id is required", http.StatusBadRequest)
	}
	err := h.permissionService.RemovePermissionFromRole(tenantID, permissionID)
	if err != nil {
		return err
	}
	return responsewriter.WriteSuccess(w, http.StatusOK, "Permission Removed Successfully!", nil)
}
