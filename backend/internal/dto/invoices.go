package dto

import "time"

type CreateInvoiceRequest struct {
	CompanyID         int64              `json:"company_id" validate:"required"`
	ClientID          int64              `json:"client_id" validate:"required"`
	BankAccountID     int64              `json:"bank_account_id" validate:"required"`
	IssueDate         string             `json:"issue_date" validate:"required"`
	DueDate           *string            `json:"due_date" validate:"omitempty"`
	Currency          string             `json:"currency" validate:"required,max=10"`
	VatRate           string             `json:"vat_rate"`
	ContractReference *string            `json:"contract_reference" validate:"omitempty,max=255"`
	ExternalReference *string            `json:"external_reference" validate:"omitempty,max=255"`
	Notes             *string            `json:"notes" validate:"omitempty,max=2000"`
	Items             []InvoiceItemInput `json:"items" validate:"required,min=1,max=10,dive"`
}

type UpdateInvoiceRequest struct {
	CompanyID         int64              `json:"company_id" validate:"required"`
	ClientID          int64              `json:"client_id" validate:"required"`
	BankAccountID     int64              `json:"bank_account_id" validate:"required"`
	InvoiceNumber     string             `json:"invoice_number" validate:"required,max=50"`
	IssueDate         string             `json:"issue_date" validate:"required"`
	DueDate           *string            `json:"due_date" validate:"omitempty"`
	Currency          string             `json:"currency" validate:"required,max=10"`
	VatRate           string             `json:"vat_rate"`
	ContractReference *string            `json:"contract_reference" validate:"omitempty,max=255"`
	ExternalReference *string            `json:"external_reference" validate:"omitempty,max=255"`
	Notes             *string            `json:"notes" validate:"omitempty,max=2000"`
	Items             []InvoiceItemInput `json:"items" validate:"required,min=1,max=10,dive"`
}

type UpdateInvoiceStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=draft sent partially_paid paid cancelled"`
}

type UpdateInvoiceOverdueRequest struct {
	IsOverdue bool `json:"is_overdue"`
}

type InvoiceItemInput struct {
	Description string `json:"description" validate:"required,max=500"`
	Quantity    string `json:"quantity" validate:"required"`
	UnitPrice   string `json:"unit_price" validate:"required"`
}

type InvoiceResponse struct {
	ID                int64                `json:"id"`
	InvoiceNumber     string               `json:"invoice_number"`
	UserID            string               `json:"user_id"`
	FamilyID          string               `json:"family_id"`
	CompanyID         int64                `json:"company_id"`
	ClientID          int64                `json:"client_id"`
	BankAccountID     int64                `json:"bank_account_id"`
	IssueDate         string               `json:"issue_date"`
	DueDate           *string              `json:"due_date"`
	Currency          string               `json:"currency"`
	Status            string               `json:"status"`
	IsOverdue         bool                 `json:"is_overdue"`
	VatRate           string               `json:"vat_rate"`
	Subtotal          string               `json:"subtotal"`
	VatAmount         string               `json:"vat_amount"`
	Total             string               `json:"total"`
	ContractReference *string              `json:"contract_reference"`
	ExternalReference *string              `json:"external_reference"`
	Notes             *string              `json:"notes"`
	Company           *CompanyResponse     `json:"company,omitempty"`
	Client            *ClientResponse      `json:"client,omitempty"`
	BankAccount       *BankAccountResponse `json:"bank_account,omitempty"`
	Items             []InvoiceItemResponse `json:"items,omitempty"`
	CreatedAt         time.Time            `json:"created_at"`
	UpdatedAt         time.Time            `json:"updated_at"`
}

type InvoiceListItem struct {
	ID            int64     `json:"id"`
	InvoiceNumber string    `json:"invoice_number"`
	IssueDate     string    `json:"issue_date"`
	DueDate       *string   `json:"due_date"`
	Status        string    `json:"status"`
	IsOverdue     bool      `json:"is_overdue"`
	Currency      string    `json:"currency"`
	CompanyName   string    `json:"company_name"`
	ClientName    string    `json:"client_name"`
	Subtotal      string    `json:"subtotal"`
	VatAmount     string    `json:"vat_amount"`
	Total         string    `json:"total"`
	ItemsCount    int64     `json:"items_count"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type InvoiceListResponse struct {
	Invoices []InvoiceListItem `json:"invoices"`
	PaginatedResponse
}

type InvoiceItemResponse struct {
	ID          int64     `json:"id"`
	InvoiceID   int64     `json:"invoice_id"`
	Description string    `json:"description"`
	Quantity    string    `json:"quantity"`
	UnitPrice   string    `json:"unit_price"`
	Total       string    `json:"total"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
