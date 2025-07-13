package models

type CreateTenantInput struct {
	Name  string  `json:"name" validate:"required"`
	Email string  `json:"email" validate:"required,email"`
	Phone *string `json:"phone"`
}

type UpdateTenantInput struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
	Phone *string `json:"phone,omitempty"`
}
