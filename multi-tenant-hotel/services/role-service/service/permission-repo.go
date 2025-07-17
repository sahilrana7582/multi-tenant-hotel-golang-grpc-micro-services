package service

import (
	"context"
	"time"

	"github.com/sahilrana7582/multi-tenant-hotel/role-service/models"
	"github.com/sahilrana7582/multi-tenant-hotel/role-service/repo"
)

type PermissionService interface {
	GivePermissionToRole(tenantID string, newPermission *models.NewPermission) (*models.Permission, error)
}

type permissionService struct {
	permissionRepo repo.PermissionRepo
}

func NewPermissionService(permissionRepo repo.PermissionRepo) PermissionService {
	return &permissionService{
		permissionRepo: permissionRepo,
	}
}

func (s *permissionService) GivePermissionToRole(tenantID string, newPermission *models.NewPermission) (*models.Permission, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.permissionRepo.GivePermissionToRole(ctx, tenantID, newPermission)
}
