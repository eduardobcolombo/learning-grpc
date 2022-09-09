package repository

import "github.com/eduardobcolombo/learning-grpc/cmd/server/domain/entity"

type PortRepository interface {
	SavePort(*entity.Port) (*entity.Port, error)
	RetrievePorts() ([]*entity.Port, error)
}
