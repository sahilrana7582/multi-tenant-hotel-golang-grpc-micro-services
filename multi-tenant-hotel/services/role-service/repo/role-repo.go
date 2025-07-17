package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sahilrana7582/multi-tenant-hotel/role-service/models"
)

type RoleRepo interface {
	NewRole(ctx context.Context, tenantID string, role *models.NewRole) error
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
