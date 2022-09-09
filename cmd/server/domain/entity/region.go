package entity

import "gorm.io/gorm"

type Region struct {
	Region string
	PortID uint
	gorm.Model
}
