package dto

import "time"

type CreateCompanyRequest struct {
	Name          string  `json:"name" validate:"required,max=255"`
	ContactPerson string  `json:"contact_person" validate:"required,max=255"`
	Address       string  `json:"address" validate:"required,max=500"`
	Phone         *string `json:"phone" validate:"omitempty,max=50"`
	VatNumber     *string `json:"vat_number" validate:"omitempty,max=50"`
	RegNumber     *string `json:"reg_number" validate:"omitempty,max=50"`
}

type UpdateCompanyRequest struct {
	Name          string  `json:"name" validate:"required,max=255"`
	ContactPerson string  `json:"contact_person" validate:"required,max=255"`
	Address       string  `json:"address" validate:"required,max=500"`
	Phone         *string `json:"phone" validate:"omitempty,max=50"`
	VatNumber     *string `json:"vat_number" validate:"omitempty,max=50"`
	RegNumber     *string `json:"reg_number" validate:"omitempty,max=50"`
}

type CompanyResponse struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	ContactPerson string    `json:"contact_person"`
	Address       string    `json:"address"`
	Phone         *string   `json:"phone"`
	VatNumber     *string   `json:"vat_number"`
	RegNumber     *string   `json:"reg_number"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
