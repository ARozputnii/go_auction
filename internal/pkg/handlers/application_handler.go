package handlers

import (
	"github.com/gorilla/mux"
	"go_auction/internal/pkg/services"
)

type ApplicationHandler struct {
	service *services.ApplicationService
}

func NewApplicationHandler(s *services.ApplicationService) *ApplicationHandler {
	return &ApplicationHandler{
		service: s,
	}
}

func (h *ApplicationHandler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	return r
}
