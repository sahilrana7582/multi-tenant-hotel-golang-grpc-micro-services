package models

type Permission struct {
	ID           string  `json:"id"`
	TenantID     string  `json:"tenant_id"`
	RoleID       string  `json:"role_id"`
	Action       string  `json:"action"`
	DepartmentID *string `json:"department_id"`
}

type PermissionByRole struct {
	RoleID          string              `json:"role_id"`
	RoleName        string              `json:"role_name"`
	RoleDescription string              `json:"role_description"`
	TenantID        string              `json:"tenant_id"`
	Permissions     []DepartmentActions `json:"permissions"`
}

type DepartmentActions struct {
	PermissionID   string  `json:"permission_id"`
	DepartmentID   *string `json:"department_id"`
	DepartmentName *string `json:"department_name"`
	Actions        string  `json:"action"`
}
