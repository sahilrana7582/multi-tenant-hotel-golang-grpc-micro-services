package models

type NewRoom struct {
	DepartmentID  string  `json:"department_id"`
	RoomID        string  `json:"room_id"`
	RoomNumber    string  `json:"room_number"`
	Floor         int     `json:"floor"`
	PricePerNight float32 `json:"price_per_night"`
	Status        string  `json:"status"`
	IsActive      bool    `json:"is_active"`
	Description   string  `json:"description"`
}

type NewRoomType struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
