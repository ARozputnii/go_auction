package config

import (
	"go_auction/internal/pkg/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Lot{})
}
