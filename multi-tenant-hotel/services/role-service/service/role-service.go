package service

import (
	"context"
	"time"

	"github.com/sahilrana7582/multi-tenant-hotel/role-service/models"
	"github.com/sahilrana7582/multi-tenant-hotel/role-service/repo"
)

type RoleService interface {
	CreateRole(tenantID string, role *models.NewRole) error
	GetRoleByID(tenantID string, roleID string) (*models.Role, error)
	GetAllRoles(tenantID string) ([]*models.Role, error)
	UpdateRole(tenantID string, roleID string, role *models.UpdateRole) error
	DeleteRole(tenantID string, roleID string) error
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

func (s *roleService) GetRoleByID(tenantID string, roleID string) (*models.Role, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.roleRepo.GetRole(ctx, tenantID, roleID)
}

func (s *roleService) GetAllRoles(tenantID string) ([]*models.Role, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return s.roleRepo.GetAllRoles(ctx, tenantID)
}

func (s *roleService) UpdateRole(tenantID string, roleID string, role *models.UpdateRole) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.roleRepo.UpdateRole(ctx, tenantID, roleID, role)
}

func (s *roleService) DeleteRole(tenantID string, roleID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.roleRepo.DeleteRole(ctx, tenantID, roleID)
}
