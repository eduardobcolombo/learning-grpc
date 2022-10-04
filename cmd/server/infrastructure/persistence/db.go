package persistence

import (
	"github.com/eduardobcolombo/learning-grpc/cmd/server/domain/repository"
	"github.com/eduardobcolombo/learning-grpc/internal/pkg/db"
)

type Repositories struct {
	Port repository.PortRepository
	DB   *db.DB
}

func New(db *db.DB) (*Repositories, error) {
	return &Repositories{
		Port: NewPortRepository(db),
		DB:   db,
	}, nil
}

func (s *Repositories) Close() error {
	return s.DB.Close()
}
