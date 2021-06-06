package interfaces

import (
	"log"

	"github.com/eduardobcolombo/learning-grpc/server/application"
	"github.com/eduardobcolombo/learning-grpc/server/domain/entity"
	"github.com/eduardobcolombo/portpb"
)

type Port struct {
	portApp application.PortAppInterface
}

func NewPort(pApp application.PortAppInterface) *Port {
	return &Port{
		portApp: pApp,
	}
}

func (p *Port) SavePort(portReq *portpb.PortRequest) error {
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
	coord.Lat = float64(portReq.GetPort().GetCoordinates().GetLat())
	coord.Long = float64(portReq.GetPort().GetCoordinates().GetLong())
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

	_, err := p.portApp.SavePort(&port)
	if err != nil {
		log.Printf("Error %v", err)
		return err
	}
	return nil
}

func (p *Port) RetrievePorts() (ports []*entity.Port, err error) {
	ports, err = p.portApp.RetrievePorts()
	if err != nil {

		log.Printf("Error retrieving data from DB: %v", err)
		return ports, err
	}
	return ports, nil
}
