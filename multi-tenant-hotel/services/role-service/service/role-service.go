package service

import (
	"context"
	"time"

	"github.com/sahilrana7582/multi-tenant-hotel/role-service/models"
	"github.com/sahilrana7582/multi-tenant-hotel/role-service/repo"
)

type RoleService interface {
	CreateRole(tenantID string, role *models.NewRole) error
}

type roleService struct {
	roleRepo repo.RoleRepo
}

func NewRoleService(roleRepo repo.RoleRepo) RoleService {
	return &roleService{
		roleRepo: roleRepo,
	}
}

func (s *roleService) CreateRole(tenantID string, role *models.NewRole) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.roleRepo.NewRole(ctx, tenantID, role)
}
