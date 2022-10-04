package db

import (
	"github.com/eduardobcolombo/learning-grpc/cmd/server/domain/entity"
	"github.com/eduardobcolombo/learning-grpc/internal/pkg/sqlDB"
)

type Store struct {
	sqlDB *sqlDB.DB
}

func NewStore(sqlDB *sqlDB.DB) Store {
	return Store{sqlDB: sqlDB}
}

// TODO: figure a way to identify the Port to allow the REPLACE/UPDATE
// instead of just add new records to the DB.
// It was not allowed in this time because the port.json data did not contains
// an unique identifier like ID or UUID or something like that
func (s *Store) Save(port *entity.Port) error {
	err := s.sqlDB.DB.Create(&port).Error
	if err != nil {
		return err
	}
	return nil
}

// TODO: Implement some limit/pagination
func (s *Store) Retrieve() (ports []*entity.Port, err error) {
	ports = []*entity.Port{}
	err = s.sqlDB.DB.Find(&ports).Error
	if err != nil {
		return nil, err
	}
	return ports, nil
}
