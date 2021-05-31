package persistence

import (
	"github.com/eduardobcolombo/learning-grpc/server/domain/entity"
	"github.com/eduardobcolombo/learning-grpc/server/domain/repository"
	"github.com/jinzhu/gorm"
)

type PortRepo struct {
	db *gorm.DB
}

func NewPortRepository(db *gorm.DB) *PortRepo {
	return &PortRepo{db}
}

var _ repository.PortRepository = &PortRepo{}

func (r *PortRepo) SavePort(port *entity.Port) (*entity.Port, map[string]string) {

	err := r.db.Debug().Create(&port).Error
	if err != nil {
		return nil, nil
	}
	return port, nil
}
