package entity

import "gorm.io/gorm"

type Port struct {
	Name        string     `json:"name"`
	City        string     `json:"city"`
	Country     string     `json:"country"`
	Alias       []Alias    `json:"alias"`
	Regions     []Region   `json:"regions"`
	Coordinates Coordinate `json:"coordinates"`
	Province    string     `json:"province"`
	Timezone    string     `json:"timezone"`
	Unlocs      []Unloc    `json:"unlocs"`
	Code        string     `json:"code"`
	gorm.Model
}
