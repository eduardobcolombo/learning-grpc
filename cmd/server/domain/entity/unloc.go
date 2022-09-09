package entity

import "gorm.io/gorm"

type Unloc struct {
	Unloc  string
	PortID uint
	gorm.Model
}
