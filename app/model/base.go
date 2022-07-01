package model

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
}

var DB *gorm.DB
