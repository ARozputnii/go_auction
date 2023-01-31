package handlers

import (
	"encoding/json"
	"go_auction/internal/pkg/models"
	"go_auction/internal/pkg/services"
	"net/http"
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
	var lot models.Lot

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&lot); err != nil {
		// RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	id, err := h.service.Create(lot)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, id)
}
