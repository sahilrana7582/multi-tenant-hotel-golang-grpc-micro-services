package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sahilrana7582/multi-tenant-hotel/department-service/models"
	"github.com/sahilrana7582/multi-tenant-hotel/department-service/service"
	"github.com/sahilrana7582/multi-tenant-hotel/pkg/errs"
)

type DepartmentHandler struct {
	departmentService service.DepartmentService
}

func NewDepartmentHandler(departmentService service.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{
		departmentService: departmentService,
	}
}

func (h *DepartmentHandler) CreateDepartment(w http.ResponseWriter, r *http.Request) error {
	var department models.DepartmentNew

	if err := json.NewDecoder(r.Body).Decode(&department); err != nil {
		return errs.New("Invalid Request Body", err.Error(), http.StatusBadRequest)
	}

	createdDepartment, err := h.departmentService.CreateDepartment(r.Context(), &department)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(createdDepartment)
}
