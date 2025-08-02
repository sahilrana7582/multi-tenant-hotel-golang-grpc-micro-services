package models

type NewHotelInfo struct {
	TenantID    string `json:"tenant_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
}

type HotelInfo struct {
	ID string `json:"id"`
	NewHotelInfo
}

type NewHotelLocation struct {
	HotelID   string  `json:"hotel_id"`
	Address   string  `json:"address"`
	City      string  `json:"city"`
	State     string  `json:"state"`
	Country   string  `json:"country"`
	ZipCode   string  `json:"zip_code"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type HotelLocationResp struct {
	Message string `json:"Message"`
}
