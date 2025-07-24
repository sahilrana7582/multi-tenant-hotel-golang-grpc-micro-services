package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sahilrana7582/multi-tenant-hotel/hotel-service/models"
	"github.com/sahilrana7582/multi-tenant-hotel/hotel-service/service"
	"github.com/sahilrana7582/multi-tenant-hotel/pkg/errs"
	responsewriter "github.com/sahilrana7582/multi-tenant-hotel/pkg/response-writer"
)

type HotelHandler struct {
	service service.IHotelService
}

func NewHotelHandler(service service.IHotelService) *HotelHandler {
	return &HotelHandler{
		service: service,
	}
}

func (h *HotelHandler) CreateHotelInfo(w http.ResponseWriter, r *http.Request) error {
	tenantID := r.Header.Get("X-Tenant-ID")
	if tenantID == "" {
		return errs.New("Missing Tenant ID", "X-Tenant-ID header is required", http.StatusBadRequest)
	}

	var newHotel models.NewHotelInfo
	if err := json.NewDecoder(r.Body).Decode(&newHotel); err != nil {
		return errs.New("Invalid request body", "Failed to decode JSON", http.StatusBadRequest)
	}
	defer r.Body.Close()

	newHotel.TenantID = tenantID
	createdHotel, err := h.service.RegisterHotelInfo(r.Context(), &newHotel)
	if err != nil {
		return err
	}

	return responsewriter.WriteSuccess(w, http.StatusCreated, "Room created successfully", createdHotel)
}
