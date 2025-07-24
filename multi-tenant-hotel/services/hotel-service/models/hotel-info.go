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
