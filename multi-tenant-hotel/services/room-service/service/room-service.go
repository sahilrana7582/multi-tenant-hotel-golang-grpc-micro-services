package service

import (
	"context"

	"github.com/sahilrana7582/multi-tenant-hotel/room-service/models"
	"github.com/sahilrana7582/multi-tenant-hotel/room-service/repo"
)

type RoomService interface {
	CreateRoom(ctx context.Context, tenantID, userID string, room *models.NewRoom) (*models.Room, error)
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
