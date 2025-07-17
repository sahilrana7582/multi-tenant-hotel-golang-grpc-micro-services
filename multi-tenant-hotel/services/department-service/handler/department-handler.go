package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sahilrana7582/multi-tenant-hotel/department-service/models"
	"github.com/sahilrana7582/multi-tenant-hotel/department-service/service"
	"github.com/sahilrana7582/multi-tenant-hotel/pkg/errs"
	responsewriter "github.com/sahilrana7582/multi-tenant-hotel/pkg/response-writer"
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

func (h *DepartmentHandler) GetDepartmentByID(w http.ResponseWriter, r *http.Request) error {
	departmentID := chi.URLParam(r, "id")
	if departmentID == "" {
		return errs.New("Invalid Request Body", "Department ID is required", http.StatusBadRequest)
	}

	tenantID := chi.URLParam(r, "tenant_id")
	if tenantID == "" {
		return errs.New("Invalid Request Body", "Tenant ID is required", http.StatusBadRequest)
	}
	department, err := h.departmentService.GetDepartmentByID(r.Context(), tenantID, departmentID)
	if err != nil {
		return err
	}

	return responsewriter.WriteSuccess(w, http.StatusOK, "Department retrieved successfully", department)
}

func (h *DepartmentHandler) GetAllDepartments(w http.ResponseWriter, r *http.Request) error {
	tenantID := chi.URLParam(r, "tenant_id")
	if tenantID == "" {
		return errs.New("Invalid Request Body", "Tenant ID is required", http.StatusBadRequest)
	}
	departments, err := h.departmentService.GetAllDepartments(r.Context(), tenantID)
	if err != nil {
		return err
	}

	return responsewriter.WriteSuccess(w, http.StatusOK, "Departments retrieved successfully", departments)
}

func (h *DepartmentHandler) UpdateDepartment(w http.ResponseWriter, r *http.Request) error {
	var department models.Department
	if err := json.NewDecoder(r.Body).Decode(&department); err != nil {
		return errs.New("Invalid Request Body", err.Error(), http.StatusBadRequest)
	}

	err := h.departmentService.UpdateDepartment(r.Context(), &department)
	if err != nil {
		return err
	}

	return responsewriter.WriteSuccess(w, http.StatusOK, "Department updated successfully", nil)

}

func (h *DepartmentHandler) DeleteDepartment(w http.ResponseWriter, r *http.Request) error {
	departmentID := chi.URLParam(r, "id")
	if departmentID == "" {
		return errs.New("Invalid Request Body", "Department ID is required", http.StatusBadRequest)
	}
	tenantID := chi.URLParam(r, "tenant_id")
	if tenantID == "" {
		return errs.New("Invalid Request Body", "Tenant ID is required", http.StatusBadRequest)
	}
	err := h.departmentService.DeleteDepartment(r.Context(), tenantID, departmentID)
	if err != nil {
		return err
	}

	return responsewriter.WriteSuccess(w, http.StatusOK, "Department deleted successfully", nil)
}
