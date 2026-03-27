package handlers

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/DigitLock/invoice-generator/backend/internal/api/middleware"
	"github.com/DigitLock/invoice-generator/backend/internal/database/sqlc"
	"github.com/DigitLock/invoice-generator/backend/internal/dto"
	"github.com/DigitLock/invoice-generator/backend/internal/repository"
)

type ClientHandler struct {
	repo *repository.ClientRepository
}

func NewClientHandler(repo *repository.ClientRepository) *ClientHandler {
	return &ClientHandler{repo: repo}
}

func (h *ClientHandler) List(w http.ResponseWriter, r *http.Request) {
	familyID, ok := middleware.GetFamilyID(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "Missing family context")
		return
	}

	status := r.URL.Query().Get("status")
	var clients []sqlc.Client
	var err error
	if status != "" {
		clients, err = h.repo.ListByStatus(r.Context(), familyID, status)
	} else {
		clients, err = h.repo.List(r.Context(), familyID)
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to list clients")
		return
	}

	resp := make([]dto.ClientResponse, 0, len(clients))
	for _, c := range clients {
		resp = append(resp, mapClient(c))
	}
	writeJSON(w, http.StatusOK, resp)
}

func (h *ClientHandler) Get(w http.ResponseWriter, r *http.Request) {
	familyID, ok := middleware.GetFamilyID(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "Missing family context")
		return
	}
	id, err := urlParamInt64(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid client ID")
		return
	}

	client, err := h.repo.GetByID(r.Context(), id, familyID)
	if err != nil {
		writeError(w, http.StatusNotFound, "Client not found")
		return
	}
	writeJSON(w, http.StatusOK, mapClient(client))
}

func (h *ClientHandler) Create(w http.ResponseWriter, r *http.Request) {
	familyID, ok := middleware.GetFamilyID(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "Missing family context")
		return
	}

	var req dto.CreateClientRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if err := validate.Struct(req); err != nil {
		writeValidationErrors(w, err)
		return
	}

	status := req.Status
	if status == "" {
		status = "active"
	}

	client, err := h.repo.Create(r.Context(), familyID, sqlc.CreateClientParams{
		Name:              req.Name,
		ContactPerson:     pgTextFromPtr(req.ContactPerson),
		Email:             pgTextFromPtr(req.Email),
		Address:           req.Address,
		VatNumber:         pgTextFromPtr(req.VatNumber),
		RegNumber:         pgTextFromPtr(req.RegNumber),
		ContractReference: pgTextFromPtr(req.ContractReference),
		ContractNotes:     pgTextFromPtr(req.ContractNotes),
		Status:            status,
	})
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to create client")
		return
	}
	writeJSON(w, http.StatusCreated, mapClient(client))
}

func (h *ClientHandler) Update(w http.ResponseWriter, r *http.Request) {
	familyID, ok := middleware.GetFamilyID(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "Missing family context")
		return
	}
	id, err := urlParamInt64(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid client ID")
		return
	}

	var req dto.UpdateClientRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if err := validate.Struct(req); err != nil {
		writeValidationErrors(w, err)
		return
	}

	client, err := h.repo.Update(r.Context(), id, familyID, sqlc.UpdateClientParams{
		Name:              req.Name,
		ContactPerson:     pgTextFromPtr(req.ContactPerson),
		Email:             pgTextFromPtr(req.Email),
		Address:           req.Address,
		VatNumber:         pgTextFromPtr(req.VatNumber),
		RegNumber:         pgTextFromPtr(req.RegNumber),
		ContractReference: pgTextFromPtr(req.ContractReference),
		ContractNotes:     pgTextFromPtr(req.ContractNotes),
		Status:            req.Status,
	})
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to update client")
		return
	}
	writeJSON(w, http.StatusOK, mapClient(client))
}

func (h *ClientHandler) Delete(w http.ResponseWriter, r *http.Request) {
	familyID, ok := middleware.GetFamilyID(r.Context())
	if !ok {
		writeError(w, http.StatusUnauthorized, "Missing family context")
		return
	}
	id, err := urlParamInt64(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid client ID")
		return
	}

	if err := h.repo.Delete(r.Context(), id, familyID); err != nil {
		writeError(w, http.StatusConflict, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func mapClient(c sqlc.Client) dto.ClientResponse {
	return dto.ClientResponse{
		ID:                c.ID,
		Name:              c.Name,
		ContactPerson:     ptrFromPgText(c.ContactPerson),
		Email:             ptrFromPgText(c.Email),
		Address:           c.Address,
		VatNumber:         ptrFromPgText(c.VatNumber),
		RegNumber:         ptrFromPgText(c.RegNumber),
		ContractReference: ptrFromPgText(c.ContractReference),
		ContractNotes:     ptrFromPgText(c.ContractNotes),
		Status:            c.Status,
		CreatedAt:         formatTime(c.CreatedAt),
		UpdatedAt:         formatTime(c.UpdatedAt),
	}
}

func pgTextFromPtr(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: *s, Valid: true}
}
