package handlers

import (
	"net/http"

	"github.com/DigitLock/invoice-generator/backend/internal/api/middleware"
	"github.com/DigitLock/invoice-generator/backend/internal/database/sqlc"
	"github.com/DigitLock/invoice-generator/backend/internal/dto"
	"github.com/DigitLock/invoice-generator/backend/internal/repository"
)

type InvoiceHandler struct {
	repo        *repository.InvoiceRepository
	companyRepo *repository.CompanyRepository
	clientRepo  *repository.ClientRepository
}

func NewInvoiceHandler(repo *repository.InvoiceRepository, companyRepo *repository.CompanyRepository, clientRepo *repository.ClientRepository) *InvoiceHandler {
	return &InvoiceHandler{repo: repo, companyRepo: companyRepo, clientRepo: clientRepo}
}

func (h *InvoiceHandler) List(w http.ResponseWriter, r *http.Request) {
	familyID, ok := middleware.GetFamilyID(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "Missing family context")
		return
	}

	page := queryInt(r, "page", 1)
	pageSize := queryInt(r, "page_size", 20)
	if pageSize > 100 {
		pageSize = 100
	}

	rows, total, err := h.repo.List(r.Context(), familyID, page, pageSize)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to list invoices")
		return
	}

	totalPages := (total + int64(pageSize) - 1) / int64(pageSize)
	items := make([]dto.InvoiceListItem, 0, len(rows))
	for _, row := range rows {
		items = append(items, dto.InvoiceListItem{
			ID:            row.ID,
			InvoiceNumber: row.InvoiceNumber,
			IssueDate:     formatDate(row.IssueDate),
			DueDate:       formatDate(row.DueDate),
			Status:        row.Status,
			IsOverdue:     row.IsOverdue,
			Currency:      row.Currency,
			CompanyName:   row.CompanyName,
			ClientName:    row.ClientName,
			Subtotal:      row.Subtotal.String(),
			VatAmount:     row.VatAmount.String(),
			Total:         row.Total.String(),
			ItemsCount:    row.ItemsCount,
			CreatedAt:     formatTime(row.CreatedAt),
			UpdatedAt:     formatTime(row.UpdatedAt),
		})
	}

	writeJSON(w, http.StatusOK, dto.InvoiceListResponse{
		Invoices: items,
		PaginatedResponse: dto.PaginatedResponse{
			Pagination: dto.PaginationMeta{
				Page:        page,
				PageSize:    pageSize,
				TotalItems:  total,
				TotalPages:  totalPages,
				HasNext:     int64(page) < totalPages,
				HasPrevious: page > 1,
			},
		},
	})
}

func (h *InvoiceHandler) Get(w http.ResponseWriter, r *http.Request) {
	familyID, ok := middleware.GetFamilyID(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "Missing family context")
		return
	}
	id, err := urlParamInt64(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid invoice ID")
		return
	}

	invoice, err := h.repo.GetByID(r.Context(), id, familyID)
	if err != nil {
		writeError(w, http.StatusNotFound, "Invoice not found")
		return
	}

	resp := mapInvoice(invoice)

	company, err := h.companyRepo.GetByID(r.Context(), invoice.CompanyID, familyID)
	if err == nil {
		c := mapCompany(company)
		resp.Company = &c
	}

	client, err := h.clientRepo.GetByID(r.Context(), invoice.ClientID, familyID)
	if err == nil {
		cl := mapClient(client)
		resp.Client = &cl
	}

	itemRows, err := h.repo.GetItems(r.Context(), id)
	if err == nil {
		resp.Items = mapInvoiceItems(itemRows)
	}

	writeJSON(w, http.StatusOK, resp)
}

func (h *InvoiceHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "Missing user context")
		return
	}
	familyID, ok := middleware.GetFamilyID(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "Missing family context")
		return
	}

	var req dto.CreateInvoiceRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if err := validate.Struct(req); err != nil {
		writeValidationErrors(w, err)
		return
	}

	if _, err := h.companyRepo.GetByID(r.Context(), req.CompanyID, familyID); err != nil {
		writeError(w, http.StatusNotFound, "Company not found")
		return
	}

	client, err := h.clientRepo.GetByID(r.Context(), req.ClientID, familyID)
	if err != nil {
		writeError(w, http.StatusNotFound, "Client not found")
		return
	}
	if client.Status != "active" {
		writeError(w, http.StatusUnprocessableEntity, "Client is inactive and cannot be used for invoicing")
		return
	}

	invoice, items, err := h.repo.Create(r.Context(), userID, familyID, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to create invoice: "+err.Error())
		return
	}

	resp := mapInvoice(invoice)
	resp.Items = mapInvoiceItems(items)
	writeJSON(w, http.StatusCreated, resp)
}

func (h *InvoiceHandler) Update(w http.ResponseWriter, r *http.Request) {
	familyID, ok := middleware.GetFamilyID(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "Missing family context")
		return
	}
	id, err := urlParamInt64(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid invoice ID")
		return
	}

	var req dto.UpdateInvoiceRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if err := validate.Struct(req); err != nil {
		writeValidationErrors(w, err)
		return
	}

	invoice, items, err := h.repo.Update(r.Context(), id, familyID, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to update invoice: "+err.Error())
		return
	}

	resp := mapInvoice(invoice)
	resp.Items = mapInvoiceItems(items)
	writeJSON(w, http.StatusOK, resp)
}

func (h *InvoiceHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	familyID, ok := middleware.GetFamilyID(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "Missing family context")
		return
	}
	id, err := urlParamInt64(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid invoice ID")
		return
	}

	var req dto.UpdateInvoiceStatusRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if err := validate.Struct(req); err != nil {
		writeValidationErrors(w, err)
		return
	}

	invoice, err := h.repo.UpdateStatus(r.Context(), id, familyID, req.Status)
	if err != nil {
		writeError(w, http.StatusConflict, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, mapInvoice(invoice))
}

func (h *InvoiceHandler) UpdateOverdue(w http.ResponseWriter, r *http.Request) {
	familyID, ok := middleware.GetFamilyID(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "Missing family context")
		return
	}
	id, err := urlParamInt64(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid invoice ID")
		return
	}

	var req dto.UpdateInvoiceOverdueRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	invoice, err := h.repo.UpdateOverdue(r.Context(), id, familyID, req.IsOverdue)
	if err != nil {
		writeError(w, http.StatusConflict, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, mapInvoice(invoice))
}

func (h *InvoiceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	familyID, ok := middleware.GetFamilyID(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "Missing family context")
		return
	}
	id, err := urlParamInt64(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid invoice ID")
		return
	}

	if err := h.repo.Delete(r.Context(), id, familyID); err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to delete invoice")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func mapInvoice(inv sqlc.Invoice) dto.InvoiceResponse {
	return dto.InvoiceResponse{
		ID:                inv.ID,
		InvoiceNumber:     inv.InvoiceNumber,
		UserID:            inv.UserID,
		FamilyID:          inv.FamilyID,
		CompanyID:         inv.CompanyID,
		ClientID:          inv.ClientID,
		BankAccountID:     inv.BankAccountID,
		IssueDate:         formatDate(inv.IssueDate),
		DueDate:           formatDate(inv.DueDate),
		Currency:          inv.Currency,
		Status:            inv.Status,
		IsOverdue:         inv.IsOverdue,
		VatRate:           inv.VatRate.String(),
		Subtotal:          inv.Subtotal.String(),
		VatAmount:         inv.VatAmount.String(),
		Total:             inv.Total.String(),
		ContractReference: ptrFromPgText(inv.ContractReference),
		ExternalReference: ptrFromPgText(inv.ExternalReference),
		Notes:             ptrFromPgText(inv.Notes),
		CreatedAt:         formatTime(inv.CreatedAt),
		UpdatedAt:         formatTime(inv.UpdatedAt),
	}
}

func mapInvoiceItems(items []sqlc.InvoiceItem) []dto.InvoiceItemResponse {
	resp := make([]dto.InvoiceItemResponse, 0, len(items))
	for _, it := range items {
		resp = append(resp, dto.InvoiceItemResponse{
			ID:          it.ID,
			InvoiceID:   it.InvoiceID,
			Description: it.Description,
			Quantity:    it.Quantity.String(),
			UnitPrice:   it.UnitPrice.String(),
			Total:       it.Total.String(),
			CreatedAt:   formatTime(it.CreatedAt),
			UpdatedAt:   formatTime(it.UpdatedAt),
		})
	}
	return resp
}
