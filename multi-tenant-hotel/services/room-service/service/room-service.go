package service

import (
	"context"

	"github.com/sahilrana7582/multi-tenant-hotel/room-service/models"
	"github.com/sahilrana7582/multi-tenant-hotel/room-service/repo"
)

type RoomService interface {
	CreateRoom(ctx context.Context, tenantID, userID string, room *models.NewRoom) (*models.Room, error)
	GetRoomByID(ctx context.Context, tenantID, userID, roomID string) (*models.RoomWithType, error)
	GetAllRooms(ctx context.Context, tenantID, userID string) ([]*models.RoomWithType, error)
}

type roomService struct {
	roomRepo repo.RoomRepo
}

func NewRoomService(roomRepo repo.RoomRepo) RoomService {
	return &roomService{
		roomRepo: roomRepo,
	}
}

func (r *roomService) CreateRoom(ctx context.Context, tenantID, userID string, room *models.NewRoom) (*models.Room, error) {
	return r.roomRepo.CreateRoom(ctx, tenantID, userID, room)
}

func (r *roomService) GetRoomByID(ctx context.Context, tenantID, userID, roomID string) (*models.RoomWithType, error) {
	return r.roomRepo.GetRoomByID(ctx, tenantID, userID, roomID)
}

func (r *roomService) GetAllRooms(ctx context.Context, tenantID, userID string) ([]*models.RoomWithType, error) {
	return r.roomRepo.GetAllRooms(ctx, tenantID, userID)
}
