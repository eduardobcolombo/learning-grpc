package repository

import "github.com/academiadaweb/learning-grpc/server/domain/entity"

type PortRepository interface {
	SavePort(*entity.Port) (*entity.Port, map[string]string)
}
