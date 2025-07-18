package models

import "time"

type Room struct {
	ID            string    `json:"id"`
	TenantID      string    `json:"tenant_id"`
	DepartmentID  string    `json:"department_id"`
	RoomID        string    `json:"room_id"`
	RoomNumber    string    `json:"room_number"`
	Floor         int       `json:"floor"`
	PricePerNight float32   `json:"price_per_night"`
	Status        string    `json:"status"`
	IsActive      bool      `json:"is_active"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type RoomType struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
