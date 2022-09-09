package entity

import "gorm.io/gorm"

type Coordinate struct {
	Lat    float64
	Long   float64
	PortID uint
	gorm.Model
}
