package service

import (
	"context"

	"github.com/sahilrana7582/multi-tenant-hotel/hotel-service/models"
	"github.com/sahilrana7582/multi-tenant-hotel/hotel-service/repo"
)

type IHotelService interface {
	RegisterHotelInfo(ctx context.Context, newHotelInfo *models.NewHotelInfo) (*models.HotelInfo, error)
	RegisterHotelAddress(ctx context.Context, newLocation *models.NewHotelLocation) (*models.HotelLocationResp, error)
	GetHotelAddress(ctx context.Context, hotelID string) (*models.NewHotelLocation, error)
}

type HotelService struct {
	repo repo.HotelRepo
}

func NewHotelRepo(repo *repo.HotelRepo) IHotelService {
	return &HotelService{
		repo: *repo,
	}
}

func (s *HotelService) RegisterHotelInfo(ctx context.Context, newHotelInfo *models.NewHotelInfo) (*models.HotelInfo, error) {
	return s.repo.CreateHotelInfo(ctx, newHotelInfo)
}

func (s *HotelService) RegisterHotelAddress(ctx context.Context, newLocation *models.NewHotelLocation) (*models.HotelLocationResp, error) {
	return s.repo.CreateLocation(ctx, newLocation)
}

func (s *HotelService) GetHotelAddress(ctx context.Context, hotelID string) (*models.NewHotelLocation, error) {
	return s.repo.GetLocation(ctx, hotelID)
}
