package handlers

import (
	"github.com/gorilla/mux"
	"go_auction/internal/pkg/services"
)

type ApplicationHandler struct {
	LotsHandler
}

func NewApplicationHandler(s *services.ApplicationService) *ApplicationHandler {
	return &ApplicationHandler{
		LotsHandler: *NewLotsHandler(s.ILotService),
	}
}

func (h *ApplicationHandler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/lots", h.LotsHandler.Create).Methods("POST")
	r.HandleFunc("/api/lots", h.LotsHandler.FindAll).Methods("GET")
	r.HandleFunc("/api/lots/{id:[0-9]+}", h.LotsHandler.FindByID).Methods("GET")
	r.HandleFunc("/api/lots/{id:[0-9]+}", h.LotsHandler.Update).Methods("PUT")
	r.HandleFunc("/api/lots/{id:[0-9]+}", h.LotsHandler.Delete).Methods("DELETE")

	return r
}
