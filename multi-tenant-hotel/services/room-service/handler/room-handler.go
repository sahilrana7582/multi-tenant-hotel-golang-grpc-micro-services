package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sahilrana7582/multi-tenant-hotel/pkg/errs"
	responsewriter "github.com/sahilrana7582/multi-tenant-hotel/pkg/response-writer"
	"github.com/sahilrana7582/multi-tenant-hotel/room-service/models"
	"github.com/sahilrana7582/multi-tenant-hotel/room-service/service"
)

type RoomHandler struct {
	roomService service.RoomService
}

func NewRoomHandler(roomService service.RoomService) *RoomHandler {
	return &RoomHandler{
		roomService: roomService,
	}
}

func (h *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) error {
	tenantID := r.Header.Get("X-Tenant-ID")

	if tenantID == "" {
		return errs.New("Invalid Tenant ID", "Tenant ID is required", http.StatusBadRequest)
	}

	userID := r.Header.Get("X-User-ID")
	if userID == "" {
		return errs.New("Invalid User ID", "User ID is required", http.StatusBadRequest)
	}

	var roomInput models.NewRoom
	if err := json.NewDecoder(r.Body).Decode(&roomInput); err != nil {
		return errs.New("Invalid Request Body", "Invalid request body", http.StatusBadRequest)
	}

	room, err := h.roomService.CreateRoom(r.Context(), tenantID, userID, &roomInput)
	if err != nil {
		return err
	}

	return responsewriter.WriteSuccess(w, http.StatusCreated, "Room created successfully", room)
}
