package service

import (
	"context"

	"github.com/sahilrana7582/hotel-mgmt/services/tenant-service/models"
	"github.com/sahilrana7582/hotel-mgmt/services/tenant-service/repo"
)

type TenantService struct {
	repo repo.TenantRepo
}

func NewTenantService(repo repo.TenantRepo) *TenantService {
	return &TenantService{repo: repo}
}

func (s *TenantService) RegisterTenant(ctx context.Context, input *models.CreateTenantInput) (*models.Tenant, error) {

	newTenant, err := s.repo.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	return newTenant, nil
}

func (s *TenantService) GetTenant(ctx context.Context, id string) (*models.Tenant, error) {
	tenant, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return tenant, nil
}

func (s *TenantService) ListTenants(ctx context.Context) ([]*models.Tenant, error) {
	tenants, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return tenants, nil
}

func (s *TenantService) UpdateTenant(ctx context.Context, id string, input *models.UpdateTenantInput) (*models.Tenant, error) {
	tenant, err := s.repo.Update(ctx, input, id)
	if err != nil {
		return nil, err
	}
	return tenant, nil
}

func (s *TenantService) DeleteTenant(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
