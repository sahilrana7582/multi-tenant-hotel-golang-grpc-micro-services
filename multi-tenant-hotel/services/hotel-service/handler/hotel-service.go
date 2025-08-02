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

func (h *HotelHandler) CreatNewLocation(w http.ResponseWriter, r *http.Request) error {
	var newLocation models.NewHotelLocation

	if err := json.NewDecoder(r.Body).Decode(&newLocation); err != nil {
		return errs.New("BAD REQUEST", "Request body is not valid", http.StatusBadRequest)
	}
	defer r.Body.Close()

	creatHotelLocation, err := h.service.RegisterHotelAddress(r.Context(), &newLocation)
	if err != nil {
		return err
	}

	return responsewriter.WriteSuccess(w, http.StatusCreated, "Location Created Successfully", creatHotelLocation)
}

func (h *HotelHandler) GetHotelLocation(w http.ResponseWriter, r *http.Request) error {
	hotelID := r.URL.Query().Get("hotel_id")
	if hotelID == "" {
		return errs.New("BAD REQUEST", "hotel_id query param is required", http.StatusBadRequest)
	}

	location, err := h.service.GetHotelAddress(r.Context(), hotelID)
	if err != nil {
		return err
	}

	if location == nil {
		return errs.New("NOT FOUND", "No location found for the given hotel_id", http.StatusNotFound)
	}

	return responsewriter.WriteSuccess(w, http.StatusOK, "Hotel location fetched successfully", location)
}
