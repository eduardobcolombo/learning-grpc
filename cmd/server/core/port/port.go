//go:generate mockgen --destination=./mock_port/mock_port.go --source=./port.go

package port

import (
	"context"

	"github.com/eduardobcolombo/learning-grpc/cmd/server/domain/entity"
)

type Storer interface {
	StorerReader
	StorerWriter
}
type StorerReader interface {
	GetByID(ctx context.Context, id uint) (port *entity.Port, err error)
	GetByUnloc(ctx context.Context, unloc string) (port *entity.Port, err error)
	GetAll(ctx context.Context) (ports []entity.Port, err error)
}
type StorerWriter interface {
	Create(ctx context.Context, port entity.Port) error
	Update(ctx context.Context, port entity.Port) error
	Delete(ctx context.Context, id uint) error
}

type Core struct {
	store Storer
}

func NewCore(store Storer) *Core {
	return &Core{
		store: store,
	}
}

func (c Core) Create(ctx context.Context, port entity.Port) error {
	return c.store.Create(ctx, port)
}

func (c Core) Update(ctx context.Context, port entity.Port) error {
	return c.store.Update(ctx, port)
}

func (c Core) GetByID(ctx context.Context, id uint) (port *entity.Port, err error) {
	return c.store.GetByID(ctx, id)
}

func (c Core) GetByUnloc(ctx context.Context, unloc string) (port *entity.Port, err error) {
	return c.store.GetByUnloc(ctx, unloc)
}

func (c Core) GetAll(ctx context.Context) ([]entity.Port, error) {
	return c.store.GetAll(ctx)
}

func (c Core) Delete(ctx context.Context, id uint) error {
	return c.store.Delete(ctx, id)
}
