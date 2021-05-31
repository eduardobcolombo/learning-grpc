package repository

import "github.com/eduardobcolombo/learning-grpc/server/domain/entity"

type PortRepository interface {
	SavePort(*entity.Port) (*entity.Port, map[string]string)
}
