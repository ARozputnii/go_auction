package services

import "go_auction/internal/pkg/models"

type ApplicationService struct {
	model *models.ApplicationModel
}

func NewApplicationService(m *models.ApplicationModel) *ApplicationService {
	return &ApplicationService{
		model: m,
	}
}
