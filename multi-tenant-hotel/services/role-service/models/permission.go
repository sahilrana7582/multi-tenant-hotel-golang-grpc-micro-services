package models

type Permission struct {
	ID           string  `json:"id"`
	TenantID     string  `json:"tenant_id"`
	RoleID       string  `json:"role_id"`
	Action       string  `json:"action"`
	DepartmentID *string `json:"department_id"`
}
