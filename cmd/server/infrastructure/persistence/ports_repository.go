package persistence

import (
	"github.com/eduardobcolombo/learning-grpc/cmd/server/domain/entity"
	"github.com/eduardobcolombo/learning-grpc/internal/pkg/db"
)

type PortRepo struct {
	db *db.DB
}

func NewPortRepository(db *db.DB) *PortRepo {
	return &PortRepo{db: db}
}

// var _ repository.PortRepository = &PortRepo{}

// TODO: figure a way to identify the Port to allow the REPLACE/UPDATE
// instead of just add new records to the DB.
// It was not allowed in this time because the port.json data did not contains
// an unique identifier like ID or UUID or something like that
func (r *PortRepo) SavePort(port *entity.Port) (*entity.Port, error) {
	err := r.db.DB.Create(&port).Error
	if err != nil {
		return nil, err
	}
	return port, nil
}

// TODO: Implement some limit/pagination
func (r *PortRepo) RetrievePorts() (ports []*entity.Port, err error) {
	ports = []*entity.Port{}
	err = r.db.DB.Find(&ports).Error
	if err != nil {
		return nil, err
	}
	return ports, nil
}
