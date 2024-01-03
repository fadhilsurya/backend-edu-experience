package models

import (
	"gorm.io/gorm"
)

type City struct {
	gorm.Model
	Name       int
	ProvinceID int
}

type Province struct {
	gorm.Model
	Name int
}
