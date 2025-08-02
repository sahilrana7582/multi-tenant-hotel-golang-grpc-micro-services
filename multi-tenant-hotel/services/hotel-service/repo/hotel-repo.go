package repo

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sahilrana7582/multi-tenant-hotel/hotel-service/models"
	"github.com/sahilrana7582/multi-tenant-hotel/pkg/errs"
)

type HotelRepo interface {
	// Hotel General Info
	CreateHotelInfo(context.Context, *models.NewHotelInfo) (*models.HotelInfo, error)

	// Hotel Locations
	CreateLocation(context.Context, *models.NewHotelLocation) (*models.HotelLocationResp, error)
	GetLocation(context.Context, string) (*models.NewHotelLocation, error)
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

func (r *HotelRepository) CreateLocation(ctx context.Context, newLoc *models.NewHotelLocation) (*models.HotelLocationResp, error) {
	query := `
		INSERT INTO hotel_locations (hotel_id, address, city, state, country, zip_code, latitude, longitude )
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
	`

	_, err := r.db.Exec(ctx, query, newLoc.HotelID, newLoc.Address, newLoc.City, newLoc.State, newLoc.Country, newLoc.ZipCode, newLoc.Latitude, newLoc.Longitude)

	if err != nil {
		return nil, errs.New("DB Error", "Something went wrong!"+err.Error(), http.StatusInternalServerError)
	}

	return &models.HotelLocationResp{
		Message: "Location created successfully",
	}, nil
}

func (r *HotelRepository) GetLocation(ctx context.Context, hotelId string) (*models.NewHotelLocation, error) {
	query := `
		SELECT hotel_id, address, city, state, country, zip_code, latitude, longitude 
		FROM hotel_locations
		WHERE hotel_id = $1;
	`

	var newLoc models.NewHotelLocation

	err := r.db.QueryRow(ctx, query, hotelId).Scan(
		&newLoc.HotelID,
		&newLoc.Address,
		&newLoc.City,
		&newLoc.State,
		&newLoc.Country,
		&newLoc.ZipCode,
		&newLoc.Latitude,
		&newLoc.Longitude,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &newLoc, nil
}
