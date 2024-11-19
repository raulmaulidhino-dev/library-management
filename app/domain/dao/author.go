package dao

import (
	"base-gin/app/domain"
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	FullName string `gorm:"size:56;not null"`
	Gender *domain.TypeGender `gorm:"type:enum('m','f');not null"`
	BirthDate *time.Time
}