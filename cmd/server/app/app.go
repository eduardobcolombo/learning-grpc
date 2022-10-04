package app

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/eduardobcolombo/learning-grpc/cmd/server/domain/entity"
	"github.com/eduardobcolombo/learning-grpc/cmd/server/handlers"
	"github.com/eduardobcolombo/learning-grpc/internal/pkg/portpb"
)

type Server struct {
	portpb.PortServiceServer
	Port handlers.Port
}

func (s *Server) Update(stream portpb.PortService_UpdateServer) error {
	ctx := context.Background()

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

		if err := s.Port.Save(ctx, req); err != nil {
			log.Printf("Error while saving port stream: %v", err)
		}

		count++
	}

}

func (s *Server) Retrieve(req *portpb.RetrievePortsRequest, stream portpb.PortService_RetrieveServer) error {
	ctx := context.Background()

	data, err := s.Port.Retrieve(ctx)
	if err != nil {
		log.Printf("Error while reading client stream: %v", err)
		return err
	}
	for _, p := range data {
		ppb := toPortpb(p)
		err = stream.Send(&portpb.RetrievePortsResponse{Port: ppb})
		if err != nil {
			log.Printf("Error while reading client stream loop: %v", err)
			return nil
		}
	}

	return nil
}

func toPortpb(port *entity.Port) *portpb.Port {
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
