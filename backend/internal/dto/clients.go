package dto

import "time"

type CreateClientRequest struct {
	Name              string  `json:"name" validate:"required,max=255"`
	ContactPerson     *string `json:"contact_person" validate:"omitempty,max=255"`
	Email             *string `json:"email" validate:"omitempty,email,max=255"`
	Address           string  `json:"address" validate:"required,max=500"`
	VatNumber         *string `json:"vat_number" validate:"omitempty,max=50"`
	RegNumber         *string `json:"reg_number" validate:"omitempty,max=50"`
	ContractReference *string `json:"contract_reference" validate:"omitempty,max=255"`
	ContractNotes     *string `json:"contract_notes" validate:"omitempty,max=2000"`
	Status            string  `json:"status" validate:"omitempty,oneof=active inactive"`
}

type UpdateClientRequest struct {
	Name              string  `json:"name" validate:"required,max=255"`
	ContactPerson     *string `json:"contact_person" validate:"omitempty,max=255"`
	Email             *string `json:"email" validate:"omitempty,email,max=255"`
	Address           string  `json:"address" validate:"required,max=500"`
	VatNumber         *string `json:"vat_number" validate:"omitempty,max=50"`
	RegNumber         *string `json:"reg_number" validate:"omitempty,max=50"`
	ContractReference *string `json:"contract_reference" validate:"omitempty,max=255"`
	ContractNotes     *string `json:"contract_notes" validate:"omitempty,max=2000"`
	Status            string  `json:"status" validate:"required,oneof=active inactive"`
}

type ClientResponse struct {
	ID                int64     `json:"id"`
	Name              string    `json:"name"`
	ContactPerson     *string   `json:"contact_person"`
	Email             *string   `json:"email"`
	Address           string    `json:"address"`
	VatNumber         *string   `json:"vat_number"`
	RegNumber         *string   `json:"reg_number"`
	ContractReference *string   `json:"contract_reference"`
	ContractNotes     *string   `json:"contract_notes"`
	Status            string    `json:"status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
