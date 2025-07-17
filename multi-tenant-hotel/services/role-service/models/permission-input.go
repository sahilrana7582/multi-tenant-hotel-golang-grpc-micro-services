package models

type NewPermission struct {
	RoleID       string  `json:"role_id"`
	Action       string  `json:"action"`
	DepartmentID *string `json:"department_id"`
}
