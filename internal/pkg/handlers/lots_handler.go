package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go_auction/internal/pkg/models"
	"go_auction/internal/pkg/services"
	"net/http"
	"strconv"
)

type LotsHandler struct {
	service services.ILotService
}

func NewLotsHandler(s services.ILotService) *LotsHandler {
	return &LotsHandler{
		service: s,
	}
}

func (h *LotsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var payload models.Lot

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	lot, err := h.service.Create(payload)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, lot)
}

func (h *LotsHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	offset, _ := strconv.Atoi(r.FormValue("offset"))

	if limit > 10 || limit < 1 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	lots, err := h.service.FindAll(limit, offset)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, lots)
}

func (h *LotsHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	lot, err := h.service.FindByID(id)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, lot)
}

func (h *LotsHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var payload models.Lot

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	lot, err := h.service.Update(id, payload)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, lot)
}

func (h *LotsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.service.Delete(id); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseMsg := fmt.Sprintf("Lot with id=%d was succesfully deleted.", id)

	RespondWithJSON(w, http.StatusOK, map[string]string{"error": responseMsg})
}
