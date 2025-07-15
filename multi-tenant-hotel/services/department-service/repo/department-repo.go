package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sahilrana7582/multi-tenant-hotel/department-service/models"
)

type DepartmentRepo interface {
	CreateDepartment(ctx context.Context, department *models.DepartmentNew) (*models.Department, error)
	// GetDepartmentByID(ctx context.Context, departmentID string) (*models.Department, error)
	// GetAllDepartments(ctx context.Context) ([]*models.Department, error)
	// UpdateDepartment(ctx context.Context, department *models.Department) error
	// DeleteDepartment(ctx context.Context, departmentID string) error
}

type departmentRepo struct {
	db *pgxpool.Pool
}

func NewDepartmentRepo(db *pgxpool.Pool) DepartmentRepo {
	return &departmentRepo{
		db: db,
	}
}

func (r *departmentRepo) CreateDepartment(ctx context.Context, department *models.DepartmentNew) (*models.Department, error) {

	query := `
		INSERT INTO departments(tenant_id, name, description)
		VALUES ($1, $2, $3)
		RETURNING id, tenant_id, name, description
	`

	var createdDepartment models.Department
	err := r.db.QueryRow(ctx, query, department.TenantID, department.Name, department.Description).Scan(&createdDepartment.ID, &createdDepartment.TenantID, &createdDepartment.Name, &createdDepartment.Description)
	if err != nil {
		return nil, err
	}
	return &createdDepartment, nil
}
