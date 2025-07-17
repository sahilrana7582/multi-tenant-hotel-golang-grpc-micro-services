package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

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
