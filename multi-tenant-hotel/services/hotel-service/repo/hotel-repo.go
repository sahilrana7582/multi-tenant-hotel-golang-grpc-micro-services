package repo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sahilrana7582/multi-tenant-hotel/hotel-service/models"
)

type HotelRepo interface {
	CreateHotelInfo(context.Context, *models.NewHotelInfo) (*models.HotelInfo, error)
}

type HotelRepository struct {
	db *pgxpool.Pool
}

func NewHotelRepo(db *pgxpool.Pool) HotelRepo {
	return &HotelRepository{
		db: db,
	}
}

func (r *HotelRepository) CreateHotelInfo(ctx context.Context, newHotelInfo *models.NewHotelInfo) (*models.HotelInfo, error) {
	var hotelInfo models.HotelInfo

	query := `
		INSERT INTO hotel_info (tenant_id, name, description, star_rating)
		VALUES ($1, $2, $3, $4)
		RETURNING id, tenant_id, name, description, star_rating
	`

	err := r.db.QueryRow(
		ctx,
		query,
		newHotelInfo.TenantID,
		newHotelInfo.Name,
		newHotelInfo.Description,
		newHotelInfo.Rating,
	).Scan(
		&hotelInfo.ID,
		&hotelInfo.TenantID,
		&hotelInfo.Name,
		&hotelInfo.Description,
		&hotelInfo.Rating,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to insert hotel info: %w", err)
	}

	return &hotelInfo, nil
}
