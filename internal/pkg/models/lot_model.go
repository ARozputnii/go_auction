package models

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type Datetime struct {
	time.Time
}

func (t *Datetime) UnmarshalJSON(input []byte) error {
	layout := "2000-01-01T00:00:00.00:00"
	strInput := strings.Trim(string(input), `"`)
	newTime, err := time.Parse(layout, strInput)
	if err != nil {
		return err
	}

	t.Time = newTime
	return nil
}

type Lot struct {
	ID             uint      `json:"id" gorm:"primary_key"`
	Title          string    `json:"title"`
	Status         string    `json:"status"`
	CurrentPrice   float32   `json:"current_price"`
	EstimatedPrice float32   `json:"estimated_price"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	// LotStartTime   time.Time `json:"lot_start_time"` // TODO should be in the future
	// LotEndTime     time.Time `json:"lot_end_time"`   // TODO should be in the future
}

type LotModel struct {
	DB *gorm.DB
}

type ILotModel interface {
	Create(payload Lot) (Lot, error)
	FindAll(limit, offset int) ([]Lot, error)
	FindByID(id int) (Lot, error)
	Update(id int, payload Lot) (Lot, error)
	Delete(id int) error
}

func NewLotModel(db *gorm.DB) *LotModel {
	return &LotModel{DB: db}
}

func (l *LotModel) Create(payload Lot) (Lot, error) {
	if result := l.DB.Create(&payload); result.Error != nil {
		return Lot{}, result.Error
	}

	return payload, nil
}

func (l *LotModel) FindAll(limit, offset int) ([]Lot, error) {
	var lots []Lot

	result := l.DB.Limit(limit).Offset(offset).Find(&lots)

	if result.Error != nil {
		return []Lot{}, result.Error
	}

	return lots, nil
}

func (l *LotModel) FindByID(id int) (Lot, error) {
	var lot Lot

	result := l.DB.First(&lot, id)

	if result.Error != nil {
		return Lot{}, result.Error
	}

	return lot, nil
}

func (l *LotModel) Update(id int, payload Lot) (Lot, error) {
	lot, err := l.FindByID(id)
	if err != nil {
		return Lot{}, err
	}

	payload.ID = lot.ID

	result := l.DB.Save(&payload)

	if result.Error != nil {
		return Lot{}, result.Error
	}

	return payload, nil
}

func (l *LotModel) Delete(id int) error {
	if _, err := l.FindByID(id); err != nil {
		return err
	}

	result := l.DB.Delete(&Lot{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
