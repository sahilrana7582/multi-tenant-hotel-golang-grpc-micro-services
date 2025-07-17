package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sahilrana7582/multi-tenant-hotel/role-service/models"
)

type PermissionRepo interface {
	GivePermissionToRole(ctx context.Context, tenantID string, newPermission *models.NewPermission) (*models.Permission, error)
}

type permissionRepo struct {
	db *pgxpool.Pool
}

func NewPermissionRepo(db *pgxpool.Pool) PermissionRepo {
	return &permissionRepo{
		db: db,
	}
}

func (r *permissionRepo) GivePermissionToRole(
	ctx context.Context,
	tenantID string,
	newPermission *models.NewPermission,
) (*models.Permission, error) {
	query := `
        INSERT INTO permissions (tenant_id, role_id, action, department_id)
        VALUES ($1, $2, $3, $4)
        RETURNING id, tenant_id, role_id, action, department_id
    `

	var permission models.Permission
	err := r.db.QueryRow(
		ctx, query,
		tenantID,
		newPermission.RoleID,
		newPermission.Action,
		newPermission.DepartmentID,
	).Scan(
		&permission.ID,
		&permission.TenantID,
		&permission.RoleID,
		&permission.Action,
		&permission.DepartmentID,
	)
	return &permission, err
}
