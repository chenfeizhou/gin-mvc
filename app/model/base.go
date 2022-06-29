package model

import (
	"gorm.io/gorm"
)

type baseModel struct {
	gorm.Model
}

var DB *gorm.DB
