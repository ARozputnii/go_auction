package models

import (
	"gorm.io/gorm"
)

type ApplicationModel struct {
	ILotModel
}

func NewApplicationModel(db *gorm.DB) *ApplicationModel {
	return &ApplicationModel{
		ILotModel: NewLotModel(db),
	}
}
