package models

type Role struct {
	TenantID    string `json:"tenant_id"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
