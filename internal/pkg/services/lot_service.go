package services

import (
	"go_auction/internal/pkg/models"
)

type LotService struct {
	lot models.ILotModel
}

type ILotService interface {
	Create(params models.Lot) (models.Lot, error)
	// FindAll() ([]any, error)
	// GetById(id int) (any, error)
	// Delete(id int) (any, error)
	// Update(id int, params any) (any, error)
}

func NewLotService(l models.ILotModel) *LotService {
	return &LotService{
		lot: l,
	}
}

func (s *LotService) Create(params models.Lot) (models.Lot, error) {
	result, err := s.lot.Create(params)

	if err != nil {
		return models.Lot{}, err
	}

	return result, nil
}
