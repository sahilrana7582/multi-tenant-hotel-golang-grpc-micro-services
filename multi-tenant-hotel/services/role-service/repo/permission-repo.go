package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sahilrana7582/multi-tenant-hotel/role-service/models"
)

type PermissionRepo interface {
	GivePermissionToRole(ctx context.Context, tenantID string, newPermission *models.NewPermission) (*models.Permission, error)
	GetPermissionsByRole(ctx context.Context, tenantID string, roleID string) (*models.PermissionByRole, error)
	GetAllRolesPermissions(ctx context.Context, tenantID string) ([]*models.PermissionByRole, error)
	RemovePermissionFromRole(ctx context.Context, tenantID string, permissionID string) error
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

func (r *permissionRepo) GetPermissionsByRole(
	ctx context.Context,
	tenantID string,
	roleID string,
) (*models.PermissionByRole, error) {
	query := `
		SELECT 
			r.id,
			r.name,
			r.description,
			r.tenant_id,
			JSON_AGG(
				JSON_BUILD_OBJECT(
					'permission_id', p.id,
					'department_id', d.id,
					'department_name', d.name,
					'action', p.action
				)
			) AS permissions
		FROM roles r
		JOIN permissions p ON p.role_id = r.id
		LEFT JOIN departments d ON d.id = p.department_id
		WHERE r.tenant_id = $1 AND r.id = $2
		GROUP BY r.id, r.name, r.description, r.tenant_id;
	`

	var permissionByRole models.PermissionByRole
	err := r.db.QueryRow(
		ctx, query,
		tenantID,
		roleID,
	).Scan(
		&permissionByRole.RoleID,
		&permissionByRole.RoleName,
		&permissionByRole.RoleDescription,
		&permissionByRole.TenantID,
		&permissionByRole.Permissions,
	)

	if err != nil {
		return nil, err
	}

	return &permissionByRole, nil
}

func (r *permissionRepo) GetAllRolesPermissions(ctx context.Context, tenantID string) ([]*models.PermissionByRole, error) {
	query := `
		SELECT
			r.id,
			r.name,
			r.description,
			r.tenant_id,
			JSON_AGG(
				JSON_BUILD_OBJECT(
					'permission_id', p.id,
					'department_id', d.id,
					'department_name', d.name,
					'action', p.action
				)
			) AS permissions
		FROM roles r
		JOIN permissions p ON p.role_id = r.id
		LEFT JOIN departments d ON d.id = p.department_id
		WHERE r.tenant_id = $1
		GROUP BY r.id, r.name, r.description, r.tenant_id;
	`

	rows, err := r.db.Query(ctx, query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissionsByRole []*models.PermissionByRole
	for rows.Next() {
		var permissionByRole models.PermissionByRole
		err := rows.Scan(
			&permissionByRole.RoleID,
			&permissionByRole.RoleName,
			&permissionByRole.RoleDescription,
			&permissionByRole.TenantID,
			&permissionByRole.Permissions,
		)
		if err != nil {
			return nil, err
		}
		permissionsByRole = append(permissionsByRole, &permissionByRole)
	}

	return permissionsByRole, nil

}

func (r *permissionRepo) RemovePermissionFromRole(
	ctx context.Context,
	tenantID string,
	permissionID string,
) error {
	query := `
		DELETE FROM permissions
		WHERE id = $1 AND tenant_id = $2
	`

	_, err := r.db.Exec(ctx, query, permissionID, tenantID)

	if err != nil {
		return err
	}

	return nil

}
