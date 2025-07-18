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
	GetRoomByID(ctx context.Context, tenantID, userID, roomID string) (*models.RoomWithType, error)
	GetAllRooms(ctx context.Context, tenantID, userID string) ([]*models.RoomWithType, error)
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

	query := `
	WITH permission_check AS (
	  SELECT check_user_permission($1, $2) AS has_permission
	)
	INSERT INTO rooms (
	  tenant_id, department_id, room_type_id,
	  room_number, floor, price_per_night, 
	  status, is_active, description
	)
	SELECT $3, $4, $5, $6, $7, $8, $9, $10, $11
	FROM permission_check
	WHERE has_permission = TRUE
	RETURNING id, tenant_id, department_id, 
	  room_type_id, room_number, floor, price_per_night,
	  status, is_active, description, created_at, updated_at;
	`

	var newRoom models.Room
	err = tx.QueryRow(ctx, query,
		userID,
		room.DepartmentID,
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

func (r *roomRepo) GetRoomByID(ctx context.Context, tenantID, userID, roomID string) (*models.RoomWithType, error) {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	query := `
		WITH permission_check AS (
			SELECT check_user_permission($1, $2) AS has_permission
		)

		SELECT 
		r.id, r.tenant_id, r.department_id,
		r.room_number, r.floor, r.price_per_night,
		r.status, r.is_active, r.description, r.created_at, r.updated_at,
		rt.id AS room_type_id, rt.name AS room_type_name,
		rt.description AS room_type_description, rt.created_at AS room_type_created_at,
		rt.updated_at AS room_type_updated_at
		FROM rooms r
		JOIN room_types rt ON r.room_type_id = rt.id
		JOIN permission_check pc ON TRUE
		WHERE r.id = $3 AND r.tenant_id = $4 AND pc.has_permission = TRUE
		ORDER BY r.room_number ASC;
	`

	var room models.RoomWithType
	err = tx.QueryRow(ctx, query, userID, tenantID, roomID, tenantID).Scan(
		&room.ID,
		&room.TenantID,
		&room.DepartmentID,
		&room.RoomNumber,
		&room.Floor,
		&room.PricePerNight,
		&room.Status,
		&room.IsActive,
		&room.Description,
		&room.CreatedAt,
		&room.UpdatedAt,
		&room.RoomType.ID,
		&room.RoomType.Name,
		&room.RoomType.Description,
		&room.RoomType.CreatedAt,
		&room.RoomType.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("get room by id: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("commit tx: %w", err)
	}

	return &room, nil
}

func (r *roomRepo) GetAllRooms(ctx context.Context, tenantID, userID string) ([]*models.RoomWithType, error) {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	query := `
		WITH permission_check AS (
			SELECT check_user_permission($1, $2) AS has_permission
		)

		SELECT 
		r.id, r.tenant_id, r.department_id,
		r.room_number, r.floor, r.price_per_night,
		r.status, r.is_active, r.description, r.created_at, r.updated_at,
		rt.id AS room_type_id, rt.name AS room_type_name,
		rt.description AS room_type_description, rt.created_at AS room_type_created_at,
		rt.updated_at AS room_type_updated_at
		FROM rooms r
		JOIN room_types rt ON r.room_type_id = rt.id
		JOIN permission_check pc ON TRUE
		WHERE r.tenant_id = $3 AND pc.has_permission = TRUE
		ORDER BY r.room_number ASC;
	`

	rows, err := tx.Query(ctx, query, userID, tenantID, tenantID)
	if err != nil {
		return nil, fmt.Errorf("get all rooms: %w", err)
	}

	defer rows.Close()

	var rooms []*models.RoomWithType
	for rows.Next() {
		var room models.RoomWithType
		err := rows.Scan(
			&room.ID,
			&room.TenantID,
			&room.DepartmentID,
			&room.RoomNumber,
			&room.Floor,
			&room.PricePerNight,
			&room.Status,
			&room.IsActive,
			&room.Description,
			&room.CreatedAt,
			&room.UpdatedAt,
			&room.RoomType.ID,
			&room.RoomType.Name,
			&room.RoomType.Description,
			&room.RoomType.CreatedAt,
			&room.RoomType.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan room: %w", err)
		}
		rooms = append(rooms, &room)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("commit tx: %w", err)
	}

	return rooms, nil
}
