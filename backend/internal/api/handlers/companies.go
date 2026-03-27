package handlers

import (
	"net/http"

	"github.com/DigitLock/invoice-generator/backend/internal/api/middleware"
	"github.com/DigitLock/invoice-generator/backend/internal/database/sqlc"
	"github.com/DigitLock/invoice-generator/backend/internal/dto"
	"github.com/DigitLock/invoice-generator/backend/internal/repository"
)

type CompanyHandler struct {
	repo *repository.CompanyRepository
}

func NewCompanyHandler(repo *repository.CompanyRepository) *CompanyHandler {
	return &CompanyHandler{repo: repo}
}

func (h *CompanyHandler) List(w http.ResponseWriter, r *http.Request) {
	familyID, ok := middleware.GetFamilyID(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "Missing family context")
		return
	}

	companies, err := h.repo.List(r.Context(), familyID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to list companies")
		return
	}

	resp := make([]dto.CompanyResponse, 0, len(companies))
	for _, c := range companies {
		resp = append(resp, mapCompany(c))
	}
	writeJSON(w, http.StatusOK, resp)
}

func (h *CompanyHandler) Get(w http.ResponseWriter, r *http.Request) {
	familyID, ok := middleware.GetFamilyID(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "Missing family context")
		return
	}
	id, err := urlParamInt64(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid company ID")
		return
	}

	company, err := h.repo.GetByID(r.Context(), id, familyID)
	if err != nil {
		writeError(w, http.StatusNotFound, "Company not found")
		return
	}
	writeJSON(w, http.StatusOK, mapCompany(company))
}

func (h *CompanyHandler) Create(w http.ResponseWriter, r *http.Request) {
	familyID, ok := middleware.GetFamilyID(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "Missing family context")
		return
	}

	var req dto.CreateCompanyRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if err := validate.Struct(req); err != nil {
		writeValidationErrors(w, err)
		return
	}

	company, err := h.repo.Create(r.Context(), familyID, req.Name, req.ContactPerson, req.Address, req.Phone, req.VatNumber, req.RegNumber)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to create company")
		return
	}
	writeJSON(w, http.StatusCreated, mapCompany(company))
}

func (h *CompanyHandler) Update(w http.ResponseWriter, r *http.Request) {
	familyID, ok := middleware.GetFamilyID(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "Missing family context")
		return
	}
	id, err := urlParamInt64(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid company ID")
		return
	}

	var req dto.UpdateCompanyRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if err := validate.Struct(req); err != nil {
		writeValidationErrors(w, err)
		return
	}

	company, err := h.repo.Update(r.Context(), id, familyID, req.Name, req.ContactPerson, req.Address, req.Phone, req.VatNumber, req.RegNumber)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to update company")
		return
	}
	writeJSON(w, http.StatusOK, mapCompany(company))
}

func (h *CompanyHandler) Delete(w http.ResponseWriter, r *http.Request) {
	familyID, ok := middleware.GetFamilyID(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "Missing family context")
		return
	}
	id, err := urlParamInt64(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid company ID")
		return
	}

	if err := h.repo.Delete(r.Context(), id, familyID); err != nil {
		writeError(w, http.StatusConflict, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func mapCompany(c sqlc.Company) dto.CompanyResponse {
	return dto.CompanyResponse{
		ID:            c.ID,
		Name:          c.Name,
		ContactPerson: c.ContactPerson,
		Address:       c.Address,
		Phone:         ptrFromPgText(c.Phone),
		VatNumber:     ptrFromPgText(c.VatNumber),
		RegNumber:     ptrFromPgText(c.RegNumber),
		CreatedAt:     formatTime(c.CreatedAt),
		UpdatedAt:     formatTime(c.UpdatedAt),
	}
}
