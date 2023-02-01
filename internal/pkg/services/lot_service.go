package services

import (
	"go_auction/internal/pkg/models"
)

type LotService struct {
	lot models.ILotModel
}

type ILotService interface {
	Create(payload models.Lot) (models.Lot, error)
	FindAll(limit, offset int) ([]models.Lot, error)
	FindByID(id int) (models.Lot, error)
	Update(id int, payload models.Lot) (models.Lot, error)
	Delete(id int) error
}

func NewLotService(l models.ILotModel) *LotService {
	return &LotService{
		lot: l,
	}
}

func (s *LotService) Create(payload models.Lot) (models.Lot, error) {
	result, err := s.lot.Create(payload)

	if err != nil {
		return models.Lot{}, err
	}

	return result, nil
}

func (s *LotService) FindAll(limit, offset int) ([]models.Lot, error) {
	result, err := s.lot.FindAll(limit, offset)

	if err != nil {
		return []models.Lot{}, err
	}

	return result, nil
}

func (s *LotService) FindByID(id int) (models.Lot, error) {
	result, err := s.lot.FindByID(id)

	if err != nil {
		return models.Lot{}, err
	}

	return result, nil
}

func (s *LotService) Update(id int, payload models.Lot) (models.Lot, error) {
	result, err := s.lot.Update(id, payload)

	if err != nil {
		return models.Lot{}, err
	}

	return result, nil
}

func (s *LotService) Delete(id int) error {
	if err := s.lot.Delete(id); err != nil {
		return err
	}

	return nil
}
