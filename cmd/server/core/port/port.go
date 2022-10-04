package port

import (
	"context"

	"github.com/eduardobcolombo/learning-grpc/cmd/server/core/port/db"
	"github.com/eduardobcolombo/learning-grpc/cmd/server/domain/entity"
	"github.com/eduardobcolombo/learning-grpc/internal/pkg/sqlDB"
)

type Core struct {
	store db.Store
}

func NewCore(sql *sqlDB.DB) Core {
	return Core{
		store: db.NewStore(sql),
	}
}

func (c Core) Save(ctx context.Context, port *entity.Port) error {
	return c.store.Save(port)
}

func (c Core) Retrieve(ctx context.Context) ([]*entity.Port, error) {
	return c.store.Retrieve()
}
