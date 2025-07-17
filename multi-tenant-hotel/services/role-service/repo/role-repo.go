package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sahilrana7582/multi-tenant-hotel/role-service/models"
)

type RoleRepo interface {
	NewRole(ctx context.Context, tenantID string, role *models.NewRole) error
	GetRole(ctx context.Context, tenantID string, roleID string) (*models.Role, error)
	GetAllRoles(ctx context.Context, tenantID string) ([]*models.Role, error)
	UpdateRole(ctx context.Context, tenantID string, roleID string, role *models.UpdateRole) error
	DeleteRole(ctx context.Context, tenantID string, roleID string) error
}

type roleRepo struct {
	db *pgxpool.Pool
}

func NewRoleRepo(db *pgxpool.Pool) RoleRepo {
	return &roleRepo{
		db: db,
	}
}

func (r *roleRepo) NewRole(ctx context.Context, tenantID string, role *models.NewRole) error {
	query := `
		INSERT INTO roles (tenant_id, name, description)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(ctx, query, tenantID, role.Name, role.Description)
	if err != nil {
		return err
	}

	return nil
}

func (r *roleRepo) GetRole(ctx context.Context, tenantID string, roleID string) (*models.Role, error) {
	query := `
		SELECT id, tenant_id, name, description
		FROM roles
		WHERE tenant_id = $1 AND id = $2
	`

	var role models.Role
	err := r.db.QueryRow(ctx, query, tenantID, roleID).Scan(&role.ID, &role.TenantID, &role.Name, &role.Description)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepo) GetAllRoles(ctx context.Context, tenantID string) ([]*models.Role, error) {
	query := `
		SELECT id, tenant_id, name, description
		FROM roles
		WHERE tenant_id = $1
	`

	rows, err := r.db.Query(ctx, query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var roles []*models.Role

	for rows.Next() {
		var role models.Role
		err := rows.Scan(&role.ID, &role.TenantID, &role.Name, &role.Description)
		if err != nil {
			return nil, err
		}
		roles = append(roles, &role)
	}

	return roles, nil
}

func (r *roleRepo) UpdateRole(ctx context.Context, tenantID string, roleID string, role *models.UpdateRole) error {
	query := `
		UPDATE roles
		SET name = $1, description = $2
		WHERE tenant_id = $3 AND id = $4
	`

	_, err := r.db.Exec(ctx, query, role.Name, role.Description, tenantID, roleID)
	if err != nil {
		return err
	}
	return nil
}

func (r *roleRepo) DeleteRole(ctx context.Context, tenantID string, roleID string) error {
	query := `
		DELETE FROM roles
		WHERE tenant_id = $1 AND id = $2
	`
	_, err := r.db.Exec(ctx, query, tenantID, roleID)
	if err != nil {
		return err
	}
	return nil
}
