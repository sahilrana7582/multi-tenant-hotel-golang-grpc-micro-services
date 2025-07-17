package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sahilrana7582/multi-tenant-hotel/pkg/auth"
	"github.com/sahilrana7582/multi-tenant-hotel/pkg/errs"
	responsewriter "github.com/sahilrana7582/multi-tenant-hotel/pkg/response-writer"
	"github.com/sahilrana7582/multi-tenant-hotel/role-service/models"
	"github.com/sahilrana7582/multi-tenant-hotel/role-service/service"
)

type RoleHandler struct {
	roleService service.RoleService
}

func NewRoleHandler(roleService service.RoleService) *RoleHandler {
	return &RoleHandler{
		roleService: roleService,
	}
}

func (h *RoleHandler) CreateRole(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("CreateRole called")
	tenantID := auth.GetTenantID(r)
	fmt.Println("Tenant ID:", tenantID)

	if tenantID == "" {

		return errs.New("No Tenant ID", "Tenant ID is required", http.StatusBadRequest)
	}

	var newRole models.NewRole
	err := json.NewDecoder(r.Body).Decode(&newRole)
	if err != nil {
		return errs.New("Invalid Request Body", "Invalid request body"+err.Error(), http.StatusBadRequest)

	}

	fmt.Println("Tenant ID:", tenantID)
	err = h.roleService.CreateRole(tenantID, &newRole)
	if err != nil {
		return errs.New("Failed to create role", "Failed to create role:"+err.Error(), http.StatusInternalServerError)
	}

	return responsewriter.WriteSuccess(w, http.StatusCreated, "Role created successfully", newRole)
}

func (h *RoleHandler) GetRoleByID(w http.ResponseWriter, r *http.Request) error {
	tenantID := auth.GetTenantID(r)
	if tenantID == "" {
		return errs.New("No Tenant ID", "Tenant ID is required", http.StatusBadRequest)
	}

	roleID := chi.URLParam(r, "id")
	if roleID == "" {
		return errs.New("Invalid Role ID", "Role ID is required", http.StatusBadRequest)
	}
	role, err := h.roleService.GetRoleByID(tenantID, roleID)
	if err != nil {
		return errs.New("Failed to get role", "Failed to get role:"+err.Error(), http.StatusInternalServerError)
	}
	return responsewriter.WriteSuccess(w, http.StatusOK, "Role fetched successfully", role)
}

func (h *RoleHandler) GetAllRoles(w http.ResponseWriter, r *http.Request) error {
	tenantID := auth.GetTenantID(r)
	if tenantID == "" {
		return errs.New("No Tenant ID", "Tenant ID is required", http.StatusBadRequest)
	}
	roles, err := h.roleService.GetAllRoles(tenantID)
	if err != nil {
		return errs.New("Failed to get roles", "Failed to get roles:"+err.Error(), http.StatusInternalServerError)
	}
	return responsewriter.WriteSuccess(w, http.StatusOK, "Roles fetched successfully", roles)
}

func (h *RoleHandler) UpdateRole(w http.ResponseWriter, r *http.Request) error {
	tenantID := auth.GetTenantID(r)
	if tenantID == "" {
		return errs.New("No Tenant ID", "Tenant ID is required", http.StatusBadRequest)
	}

	roleID := chi.URLParam(r, "id")
	if roleID == "" {
		return errs.New("Invalid Role ID", "Role ID is required", http.StatusBadRequest)
	}

	var updateRole models.UpdateRole
	err := json.NewDecoder(r.Body).Decode(&updateRole)
	if err != nil {
		return errs.New("Invalid Request Body", "Invalid request body"+err.Error(), http.StatusBadRequest)
	}

	err = h.roleService.UpdateRole(tenantID, roleID, &updateRole)
	if err != nil {
		return errs.New("Failed to update role", "Failed to update role:"+err.Error(), http.StatusInternalServerError)
	}

	return responsewriter.WriteSuccess(w, http.StatusOK, "Role updated successfully", map[string]string{"message": "Role updated successfully"})
}

func (h *RoleHandler) DeleteRole(w http.ResponseWriter, r *http.Request) error {
	tenantID := auth.GetTenantID(r)
	if tenantID == "" {
		return errs.New("No Tenant ID", "Tenant ID is required", http.StatusBadRequest)
	}
	roleID := chi.URLParam(r, "id")
	if roleID == "" {
		return errs.New("Invalid Role ID", "Role ID is required", http.StatusBadRequest)
	}
	err := h.roleService.DeleteRole(tenantID, roleID)
	if err != nil {
		return errs.New("Failed to delete role", "Failed to delete role:"+err.Error(), http.StatusInternalServerError)
	}
	return responsewriter.WriteSuccess(w, http.StatusOK, "Role deleted successfully", map[string]string{"message": "Role deleted successfully"})
}
