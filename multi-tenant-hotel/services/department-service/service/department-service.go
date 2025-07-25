package service

import (
	"context"

	"github.com/sahilrana7582/multi-tenant-hotel/department-service/models"
	"github.com/sahilrana7582/multi-tenant-hotel/department-service/repo"
)

type DepartmentService interface {
	CreateDepartment(ctx context.Context, tenandID string, department *models.DepartmentNew) (*models.Department, error)
	GetDepartmentByID(ctx context.Context, tenantID, userId, departmentID string) (*models.Department, error)
	GetAllDepartments(ctx context.Context, tenantID, userID string) ([]*models.Department, error)
	UpdateDepartment(ctx context.Context, tenantID, userID, departmentID string, department *models.DepartmentUpdate) error
	DeleteDepartment(ctx context.Context, tenantID, userID, departmentID string) error
}

type departmentService struct {
	repo repo.DepartmentRepo
}

func NewDepartmentService(repo repo.DepartmentRepo) DepartmentService {
	return &departmentService{
		repo: repo,
	}
}

func (s *departmentService) CreateDepartment(ctx context.Context, tenandID string, department *models.DepartmentNew) (*models.Department, error) {
	return s.repo.CreateDepartment(ctx, tenandID, department)
}

func (s *departmentService) GetDepartmentByID(ctx context.Context, tenantID, userID, departmentID string) (*models.Department, error) {
	return s.repo.GetDepartmentByID(ctx, tenantID, userID, departmentID)
}

func (s *departmentService) GetAllDepartments(ctx context.Context, tenantID, userID string) ([]*models.Department, error) {
	return s.repo.GetAllDepartments(ctx, tenantID, userID)
}

func (s *departmentService) UpdateDepartment(ctx context.Context, tenantID, userID, departmentID string, department *models.DepartmentUpdate) error {
	return s.repo.UpdateDepartment(ctx, tenantID, userID, departmentID, department)
}

func (s *departmentService) DeleteDepartment(ctx context.Context, tenantID, userID, departmentID string) error {
	return s.repo.DeleteDepartment(ctx, tenantID, userID, departmentID)
}
