package models

type DepartmentNew struct {
	TenantID    string `json:"tenant_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
