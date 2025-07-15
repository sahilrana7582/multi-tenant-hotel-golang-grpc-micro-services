package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sahilrana7582/multi-tenant-hotel/department-service/models"
)

type DepartmentRepo interface {
	CreateDepartment(ctx context.Context, department *models.DepartmentNew) (*models.Department, error)
	GetDepartmentByID(ctx context.Context, tenantID, departmentID string) (*models.Department, error)
	GetAllDepartments(ctx context.Context, tenantID string) ([]*models.Department, error)
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

func (r *departmentRepo) GetDepartmentByID(ctx context.Context, tenantID, departmentID string) (*models.Department, error) {
	query := `
		SELECT id, tenant_id, name, description
		FROM departments
		WHERE id = $1
	`
	var department models.Department
	err := r.db.QueryRow(ctx, query, departmentID).Scan(&department.ID, &department.TenantID, &department.Name, &department.Description)
	if err != nil {
		return nil, err
	}
	return &department, nil
}

func (r *departmentRepo) GetAllDepartments(ctx context.Context, tenantID string) ([]*models.Department, error) {
	query := `
		SELECT id, tenant_id, name, description
		FROM departments
		WHERE tenant_id = $1
	`
	rows, err := r.db.Query(ctx, query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []*models.Department

	for rows.Next() {
		var department models.Department

		if err := rows.Scan(&department.ID, &department.TenantID, &department.Name, &department.Description); err != nil {
			return nil, err
		}
		departments = append(departments, &department)
	}

	return departments, nil
}
