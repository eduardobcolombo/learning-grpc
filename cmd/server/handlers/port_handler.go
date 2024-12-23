package handlers

import (
	"context"
	"errors"

	"github.com/eduardobcolombo/learning-grpc/cmd/server/core/port"
	"github.com/eduardobcolombo/learning-grpc/cmd/server/domain/entity"
	"github.com/eduardobcolombo/learning-grpc/portpb"
)

type Port struct {
	core *port.Core
}

func NewPort(core *port.Core) *Port {
	return &Port{
		core: core,
	}
}

func (p *Port) Create(ctx context.Context, portReq *portpb.PortRequest) error {
	port := entity.Port{
		Name:      portReq.GetPort().GetName(),
		City:      portReq.GetPort().GetCity(),
		Country:   portReq.GetPort().GetCountry(),
		Alias:     portReq.GetPort().GetAlias(),
		Regions:   portReq.GetPort().GetRegions(),
		Latitude:  portReq.GetPort().GetLatitude(),
		Longitude: portReq.GetPort().GetLongitude(),
		Province:  portReq.GetPort().GetProvince(),
		Timezone:  portReq.GetPort().GetTimezone(),
		Unlocs:    portReq.GetPort().GetUnlocs(),
		Code:      portReq.GetPort().GetCode(),
	}

	if err := p.core.Create(ctx, port); err != nil {
		return err
	}

	return nil
}

func (p *Port) Update(ctx context.Context, portReq *portpb.PortRequest) error {
	port := entity.Port{
		Name:      portReq.GetPort().GetName(),
		City:      portReq.GetPort().GetCity(),
		Country:   portReq.GetPort().GetCountry(),
		Alias:     portReq.GetPort().GetAlias(),
		Regions:   portReq.GetPort().GetRegions(),
		Latitude:  portReq.GetPort().GetLatitude(),
		Longitude: portReq.GetPort().GetLongitude(),
		Province:  portReq.GetPort().GetProvince(),
		Timezone:  portReq.GetPort().GetTimezone(),
		Unlocs:    portReq.GetPort().GetUnlocs(),
		Code:      portReq.GetPort().GetCode(),
	}

	// check if the resource exists
	data, err := p.GetByUnloc(ctx, port.Unlocs)
	if err != nil {
		if errors.Is(err, entity.ErrPortNotFound) {
			// if it does not exists, try to create it
			return p.core.Create(ctx, port)
		}
		return err
	}
	// then update it
	port.ID = data.ID
	if err := p.core.Update(ctx, port); err != nil {
		return err
	}

	return nil
}

func (p *Port) GetByID(ctx context.Context, id uint) (*entity.Port, error) {
	port, err := p.core.GetByID(ctx, id)
	if err != nil {
		return port, err
	}

	return port, nil
}

func (p *Port) GetByUnloc(ctx context.Context, unloc string) (*entity.Port, error) {
	port, err := p.core.GetByUnloc(ctx, unloc)
	if err != nil {
		return port, err
	}

	return port, nil
}

func (p *Port) GetAll(ctx context.Context) (ports []entity.Port, err error) {
	ports, err = p.core.GetAll(ctx)
	if err != nil {
		return ports, err
	}

	return ports, nil
}

func (p *Port) Delete(ctx context.Context, portReq *portpb.PortIdRequest) error {

	id := uint(portReq.GetId())

	if err := p.core.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
