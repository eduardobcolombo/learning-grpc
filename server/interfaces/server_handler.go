package interfaces

import (
	"fmt"
	"io"
	"log"

	"github.com/eduardobcolombo/learning-grpc/server/domain/entity"
	"github.com/eduardobcolombo/learning-grpc/server/infrastructure/persistence"
	"github.com/eduardobcolombo/portpb"
)

type Server struct {
	portpb.PortServiceServer
	Services *persistence.Repositories
}

func (s *Server) PortsUpdate(stream portpb.PortService_PortsUpdateServer) error {
	ports := NewPort(s.Services.Port)

	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&portpb.PortResponse{
				Result: fmt.Sprintf("Received %d records.", count),
			})
		}

		if err != nil {
			log.Printf("Error while reading client stream: %v", err)
		}

		ports.SavePort(req)

		count++
	}

}

func (s *Server) PortsList(req *portpb.ListPortsRequest, stream portpb.PortService_PortsListServer) error {
	ports := NewPort(s.Services.Port)

	lPorts, err := ports.RetrievePorts()
	if err != nil {
		log.Printf("Error while reading client stream: %v", err)
		return err
	}
	for _, p := range lPorts {
		ppb := fillPortpbWithPort(p)
		err = stream.Send(&portpb.ListPortsResponse{Port: ppb})
		if err != nil {
			log.Printf("Error while reading client stream loop: %v", err)
			return nil
		}
	}

	return nil
}

func fillPortpbWithPort(port *entity.Port) *portpb.Port {
	id := int32(port.ID)
	return &portpb.Port{
		Id:          id,
		Name:        port.Name,
		City:        port.City,
		Country:     port.Country,
		Alias:       []string{},
		Regions:     []string{},
		Coordinates: &portpb.Coordinates{},
		Province:    port.Province,
		Timezone:    port.Timezone,
		Unlocs:      &portpb.Unlocs{},
		Code:        port.Code,
	}
}
