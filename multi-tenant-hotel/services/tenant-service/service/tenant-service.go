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
