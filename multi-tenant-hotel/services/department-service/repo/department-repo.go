package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sahilrana7582/multi-tenant-hotel/department-service/models"
)

type DepartmentRepo interface {
	CreateDepartment(ctx context.Context, tenandID string, department *models.DepartmentNew) (*models.Department, error)
	GetDepartmentByID(ctx context.Context, tenantID, userID, departmentID string) (*models.Department, error)
	GetAllDepartments(ctx context.Context, tenantID, userID string) ([]*models.Department, error)
	UpdateDepartment(ctx context.Context, tenandID, userID, departmentID string, department *models.DepartmentUpdate) error
	DeleteDepartment(ctx context.Context, tenantId, userId, departmentID string) error
}

type departmentRepo struct {
	db *pgxpool.Pool
}

func NewDepartmentRepo(db *pgxpool.Pool) DepartmentRepo {
	return &departmentRepo{
		db: db,
	}
}

func (r *departmentRepo) CreateDepartment(ctx context.Context, tenandID string, department *models.DepartmentNew) (*models.Department, error) {

	query := `
		INSERT INTO departments(tenant_id, name, description)
		VALUES ($1, $2, $3)
		RETURNING id, tenant_id, name, description
	`

	var createdDepartment models.Department
	err := r.db.QueryRow(ctx, query, tenandID, department.Name, department.Description).Scan(&createdDepartment.ID, &createdDepartment.TenantID, &createdDepartment.Name, &createdDepartment.Description)
	if err != nil {
		return nil, err
	}
	return &createdDepartment, nil
}

func (r *departmentRepo) GetDepartmentByID(ctx context.Context, tenantID, userId, departmentID string) (*models.Department, error) {
	query := `
		SELECT d.id, d.tenant_id, d.name, d.description
		FROM departments d
		JOIN permissions p ON p.department_id = d.id
		JOIN roles r ON r.id = p.role_id
		JOIN public.user_roles ur ON ur.role_id = p.role_id
		WHERE d.id = $1 AND d.tenant_id = $2 AND ur.user_id = $3
	`
	var department models.Department
	err := r.db.QueryRow(ctx, query, departmentID, tenantID, userId).Scan(&department.ID, &department.TenantID, &department.Name, &department.Description)
	if err != nil {
		return nil, err
	}
	return &department, nil
}

func (r *departmentRepo) GetAllDepartments(ctx context.Context, tenantID, userID string) ([]*models.Department, error) {
	query := `
		SELECT d.id, d.tenant_id, d.name, d.description
		FROM departments d
		JOIN permissions p ON p.department_id = d.id
		JOIN roles r ON r.id = p.role_id
		JOIN public.user_roles ur ON ur.role_id = p.role_id
		WHERE d.tenant_id = $1 AND ur.user_id = $2
	`
	rows, err := r.db.Query(ctx, query, tenantID, userID)
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

func (r *departmentRepo) UpdateDepartment(ctx context.Context, tenantID, userID, departmentID string, department *models.DepartmentUpdate) error {
	query := `
		UPDATE departments AS d
		SET
			name = $1,
			description = $2
		FROM permissions AS p
		INNER JOIN roles AS r ON r.id = p.role_id
		INNER JOIN user_roles AS ur ON ur.role_id = p.role_id
		WHERE
			d.id = $3
			AND d.tenant_id = $4
			AND p.department_id = d.id
			AND ur.user_id = $5
	`

	_, err := r.db.Exec(ctx, query, department.Name, department.Description, departmentID, tenantID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (r *departmentRepo) DeleteDepartment(ctx context.Context, tenantId, userId, departmentID string) error {
	query := `
		DELETE FROM departments AS d
		USING permissions AS p
		JOIN roles AS r ON r.id = p.role_id
		JOIN public.user_roles AS ur ON ur.role_id = p.role_id
		WHERE
		d.id = $1
		AND d.tenant_id = $2
		AND p.department_id = d.id
		AND ur.user_id = $3
	`
	_, err := r.db.Exec(ctx, query, departmentID, tenantId, userId)
	if err != nil {
		return err
	}
	return nil
}
