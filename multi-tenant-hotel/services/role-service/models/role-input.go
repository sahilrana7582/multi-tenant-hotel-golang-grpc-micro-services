package models

type NewRole struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateRole struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
