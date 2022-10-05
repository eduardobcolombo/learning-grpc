package handlers

import (
	"context"
	"log"

	"github.com/eduardobcolombo/learning-grpc/cmd/server/core/port"
	"github.com/eduardobcolombo/learning-grpc/cmd/server/domain/entity"
	"github.com/eduardobcolombo/learning-grpc/internal/pkg/portpb"
)

type Port struct {
	core port.Core
}

func NewPort(core port.Core) Port {
	return Port{
		core: core,
	}
}

func (p *Port) Save(ctx context.Context, portReq *portpb.PortRequest) error {
	port := entity.Port{}
	port.City = portReq.GetPort().GetCity()
	port.Country = portReq.GetPort().GetCountry()
	alias := []entity.Alias{}
	aliasReq := portReq.GetPort().GetAlias()
	for _, al := range aliasReq {
		alias = append(alias, entity.Alias{Alias: al})
	}
	port.Alias = alias

	regions := []entity.Region{}
	regionsReq := portReq.GetPort().GetRegions()
	for _, reg := range regionsReq {
		regions = append(regions, entity.Region{Region: reg})
	}
	port.Regions = regions

	coord := entity.Coordinate{}
	coord.Lat = portReq.GetPort().GetCoordinates().GetLat()
	coord.Long = portReq.GetPort().GetCoordinates().GetLong()
	port.Coordinates = coord

	port.Province = portReq.GetPort().GetProvince()
	port.Timezone = portReq.GetPort().GetTimezone()

	unlocs := []entity.Unloc{}
	unlocsReq := portReq.GetPort().GetUnlocs().GetUnloc()
	for _, unl := range unlocsReq {
		unlocs = append(unlocs, entity.Unloc{Unloc: unl})
	}
	port.Unlocs = unlocs

	port.Code = portReq.GetPort().GetCode()

	if err := p.core.Save(ctx, &port); err != nil {
		log.Printf("Error %v", err)
		return err
	}
	return nil
}

func (p *Port) Retrieve(ctx context.Context) (ports []*entity.Port, err error) {
	ports, err = p.core.Retrieve(ctx)
	if err != nil {

		log.Printf("Error retrieving data: %v", err)
		return ports, err
	}
	return ports, nil
}
