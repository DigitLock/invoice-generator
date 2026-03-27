package handlers

import (
	"net/http"

	"github.com/DigitLock/invoice-generator/backend/internal/api/middleware"
	"github.com/DigitLock/invoice-generator/backend/internal/database/sqlc"
	"github.com/DigitLock/invoice-generator/backend/internal/dto"
	"github.com/DigitLock/invoice-generator/backend/internal/repository"
)

type BankAccountHandler struct {
	repo        *repository.BankAccountRepository
	companyRepo *repository.CompanyRepository
}

func NewBankAccountHandler(repo *repository.BankAccountRepository, companyRepo *repository.CompanyRepository) *BankAccountHandler {
	return &BankAccountHandler{repo: repo, companyRepo: companyRepo}
}

func (h *BankAccountHandler) List(w http.ResponseWriter, r *http.Request) {
	familyID, ok := middleware.GetFamilyID(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "Missing family context")
		return
	}
	companyID, err := urlParamInt64(r, "company_id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid company ID")
		return
	}

	if _, err := h.companyRepo.GetByID(r.Context(), companyID, familyID); err != nil {
		writeError(w, http.StatusNotFound, "Company not found")
		return
	}

	accounts, err := h.repo.ListByCompany(r.Context(), companyID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to list bank accounts")
		return
	}

	resp := make([]dto.BankAccountResponse, 0, len(accounts))
	for _, a := range accounts {
		resp = append(resp, mapBankAccount(a))
	}
	writeJSON(w, http.StatusOK, resp)
}

func (h *BankAccountHandler) Create(w http.ResponseWriter, r *http.Request) {
	familyID, ok := middleware.GetFamilyID(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "Missing family context")
		return
	}
	companyID, err := urlParamInt64(r, "company_id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid company ID")
		return
	}

	if _, err := h.companyRepo.GetByID(r.Context(), companyID, familyID); err != nil {
		writeError(w, http.StatusNotFound, "Company not found")
		return
	}

	var req dto.CreateBankAccountRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if err := validate.Struct(req); err != nil {
		writeValidationErrors(w, err)
		return
	}

	account, err := h.repo.Create(r.Context(), sqlc.CreateBankAccountParams{
		CompanyID:     companyID,
		BankName:      req.BankName,
		BankAddress:   req.BankAddress,
		AccountHolder: req.AccountHolder,
		Iban:          req.IBAN,
		Swift:         req.SWIFT,
		Currency:      req.Currency,
		IsDefault:     req.IsDefault,
	})
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to create bank account")
		return
	}
	writeJSON(w, http.StatusCreated, mapBankAccount(account))
}

func (h *BankAccountHandler) Update(w http.ResponseWriter, r *http.Request) {
	familyID, ok := middleware.GetFamilyID(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "Missing family context")
		return
	}
	id, err := urlParamInt64(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid bank account ID")
		return
	}

	companyID, err := h.repo.GetCompanyID(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusNotFound, "Bank account not found")
		return
	}
	if _, err := h.companyRepo.GetByID(r.Context(), companyID, familyID); err != nil {
		writeError(w, http.StatusForbidden, "Access denied")
		return
	}

	var req dto.UpdateBankAccountRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if err := validate.Struct(req); err != nil {
		writeValidationErrors(w, err)
		return
	}

	account, err := h.repo.Update(r.Context(), id, sqlc.UpdateBankAccountParams{
		BankName:      req.BankName,
		BankAddress:   req.BankAddress,
		AccountHolder: req.AccountHolder,
		Iban:          req.IBAN,
		Swift:         req.SWIFT,
		Currency:      req.Currency,
		IsDefault:     req.IsDefault,
	})
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to update bank account")
		return
	}
	writeJSON(w, http.StatusOK, mapBankAccount(account))
}

func (h *BankAccountHandler) Delete(w http.ResponseWriter, r *http.Request) {
	familyID, ok := middleware.GetFamilyID(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "Missing family context")
		return
	}
	id, err := urlParamInt64(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid bank account ID")
		return
	}

	companyID, err := h.repo.GetCompanyID(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusNotFound, "Bank account not found")
		return
	}
	if _, err := h.companyRepo.GetByID(r.Context(), companyID, familyID); err != nil {
		writeError(w, http.StatusForbidden, "Access denied")
		return
	}

	if err := h.repo.Delete(r.Context(), id); err != nil {
		writeError(w, http.StatusConflict, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func mapBankAccount(a sqlc.BankAccount) dto.BankAccountResponse {
	return dto.BankAccountResponse{
		ID:            a.ID,
		CompanyID:     a.CompanyID,
		BankName:      a.BankName,
		BankAddress:   a.BankAddress,
		AccountHolder: a.AccountHolder,
		IBAN:          a.Iban,
		SWIFT:         a.Swift,
		Currency:      a.Currency,
		IsDefault:     a.IsDefault,
		CreatedAt:     formatTime(a.CreatedAt),
		UpdatedAt:     formatTime(a.UpdatedAt),
	}
}
