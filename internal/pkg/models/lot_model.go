package models

import (
	"gorm.io/gorm"
)

type Lot struct {
	ID             uint   `json:"id" gorm:"primary_key"`
	Title          string `json:"title"`
	Status         string `json:"status"`
	CurrentPrice   int32  `json:"current_price"`
	EstimatedPrice int32  `json:"estimated_price"`
	// LotStartTime   time.Time
	// LotEndTime     time.Time
	// CreatedAt      time.Time
	// UpdatedAt      time.Time
}

type LotModel struct {
	DB *gorm.DB
}

type ILotModel interface {
	Create(params Lot) (Lot, error)
}

func NewLotModel(db *gorm.DB) *LotModel {
	return &LotModel{DB: db}
}

func (l *LotModel) Create(lot Lot) (Lot, error) {
	if result := l.DB.Create(&lot); result.Error != nil {
		return Lot{}, result.Error
	}

	return lot, nil
}
