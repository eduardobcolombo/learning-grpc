package entity

import "gorm.io/gorm"

type Alias struct {
	Alias  string
	PortID uint
	gorm.Model
}
