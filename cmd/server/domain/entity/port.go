package entity

import (
	"errors"
	"time"
)

var (
	ErrPortNotFound = errors.New("port not found")
	ErrOnInsert     = errors.New("error on insert")
	ErrOnUpdate     = errors.New("error on update")
)

type Port struct {
	ID        uint
	Name      string
	City      string
	Country   string
	Alias     string
	Regions   string
	Latitude  float64
	Longitude float64
	Province  string
	Timezone  string
	Unlocs    string
	Code      string

	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
