package dto

import "time"

type CreateBankAccountRequest struct {
	BankName      string `json:"bank_name" validate:"required,max=255"`
	BankAddress   string `json:"bank_address" validate:"required,max=500"`
	AccountHolder string `json:"account_holder" validate:"required,max=255"`
	IBAN          string `json:"iban" validate:"required,min=15,max=34"`
	SWIFT         string `json:"swift" validate:"required,len=8|len=11"`
	Currency      string `json:"currency" validate:"required,max=3"`
	IsDefault     bool   `json:"is_default"`
}

type UpdateBankAccountRequest struct {
	BankName      string `json:"bank_name" validate:"required,max=255"`
	BankAddress   string `json:"bank_address" validate:"required,max=500"`
	AccountHolder string `json:"account_holder" validate:"required,max=255"`
	IBAN          string `json:"iban" validate:"required,min=15,max=34"`
	SWIFT         string `json:"swift" validate:"required,len=8|len=11"`
	Currency      string `json:"currency" validate:"required,max=3"`
	IsDefault     bool   `json:"is_default"`
}

type BankAccountResponse struct {
	ID            int64     `json:"id"`
	CompanyID     int64     `json:"company_id"`
	BankName      string    `json:"bank_name"`
	BankAddress   string    `json:"bank_address"`
	AccountHolder string    `json:"account_holder"`
	IBAN          string    `json:"iban"`
	SWIFT         string    `json:"swift"`
	Currency      string    `json:"currency"`
	IsDefault     bool      `json:"is_default"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
