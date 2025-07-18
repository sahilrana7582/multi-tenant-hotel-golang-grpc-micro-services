package repo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sahilrana7582/multi-tenant-hotel/room-service/models"
)

type RoomRepo interface {
	CreateRoom(ctx context.Context, tenantID, userID string, room *models.NewRoom) (*models.Room, error)
}

type roomRepo struct {
	db *pgxpool.Pool
}

func NewRoomRepo(db *pgxpool.Pool) RoomRepo {
	return &roomRepo{
		db: db,
	}
}

func (r *roomRepo) CreateRoom(ctx context.Context, tenantID, userID string, room *models.NewRoom) (*models.Room, error) {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `SET LOCAL app.current_user_id = $1`, userID)
	if err != nil {
		return nil, fmt.Errorf("set current_user_id: %w", err)
	}

	query := `
		INSERT INTO rooms (
			tenant_id, department_id, room_type_id,
			room_number, floor, price_per_night, 
			status, is_active, description
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, tenant_id, department_id, 
		room_type_id, room_number, floor, price_per_night,
		status, is_active, description, created_at, updated_at
	`

	var newRoom models.Room
	err = tx.QueryRow(ctx, query,
		tenantID,
		room.DepartmentID,
		room.RoomID,
		room.RoomNumber,
		room.Floor,
		room.PricePerNight,
		room.Status,
		room.IsActive,
		room.Description,
	).Scan(
		&newRoom.ID,
		&newRoom.TenantID,
		&newRoom.DepartmentID,
		&newRoom.RoomID,
		&newRoom.RoomNumber,
		&newRoom.Floor,
		&newRoom.PricePerNight,
		&newRoom.Status,
		&newRoom.IsActive,
		&newRoom.Description,
		&newRoom.CreatedAt,
		&newRoom.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("insert room: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("commit tx: %w", err)
	}

	return &newRoom, nil
}
