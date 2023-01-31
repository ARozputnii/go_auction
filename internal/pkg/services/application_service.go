package services

import (
	"go_auction/internal/pkg/models"
)

type ApplicationService struct {
	ILotService
}

func NewApplicationService(model *models.ApplicationModel) *ApplicationService {
	return &ApplicationService{
		ILotService: NewLotService(model.ILotModel),
	}
}
