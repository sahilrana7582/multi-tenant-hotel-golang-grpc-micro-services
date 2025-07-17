package models

type DepartmentNew struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type DepartmentUpdate struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
