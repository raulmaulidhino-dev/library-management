package dao

import "gorm.io/gorm"

type Publisher struct {
	gorm.Model
	Name string `gorm:"size:48;not null;unique;"`
	City string `gorm:"size:32;not null;"`
}