package persistence

import (
	"github.com/eduardobcolombo/learning-grpc/server/domain/entity"
	"github.com/eduardobcolombo/learning-grpc/server/domain/repository"
	"gorm.io/gorm"
)

type PortRepo struct {
	db *gorm.DB
}

func NewPortRepository(db *gorm.DB) *PortRepo {
	return &PortRepo{db}
}

var _ repository.PortRepository = &PortRepo{}

// TODO: figure a way to identify the Port to allow the REPLACE/UPDATE
// instead of just add new records to the DB.
// It was not allowed in this time because the port.json data did not contains
// an unique identifier like ID or UUID or something like that
func (r *PortRepo) SavePort(port *entity.Port) (*entity.Port, error) {
	err := r.db.Create(&port).Error
	if err != nil {
		return nil, err
	}
	return port, nil
}

// TODO: Implement some limit/pagination
func (r *PortRepo) RetrievePorts() (ports []*entity.Port, err error) {
	ports = []*entity.Port{}
	err = r.db.Find(&ports).Error
	if err != nil {
		return nil, err
	}
	return ports, nil
}
