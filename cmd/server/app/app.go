package app

import (
	"context"
	"fmt"
	"io"

	"github.com/eduardobcolombo/learning-grpc/cmd/server/domain/entity"
	"github.com/eduardobcolombo/learning-grpc/cmd/server/handlers"
	"github.com/eduardobcolombo/learning-grpc/foundation"
	"github.com/eduardobcolombo/learning-grpc/portpb"
)

type Server struct {
	portpb.PortServiceServer
	Port handlers.Port
	Log  *foundation.Logger
}

func (s *Server) UpdateAll(stream portpb.PortService_UpdateAllServer) error {
	ctx := context.Background()

	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&portpb.PortStringResponse{
				Result: fmt.Sprintf("Received %d records.", count),
			})
		}

		s.Log.Infof("Received record %d", count)
		if err != nil {
			s.Log.Errorf("Error while reading client stream: %v", err)
		}

		if err := s.Port.Update(ctx, req); err != nil {
			s.Log.Errorf("Error while saving port stream: %v", err)
		}

		count++
	}

}

func (s *Server) GetAll(ga *portpb.PortsGetAllRequest, stream portpb.PortService_GetAllServer) error {
	ctx := context.Background()

	data, err := s.Port.GetAll(ctx)
	if err != nil {
		s.Log.Errorf("error while reading client stream: %v", err)
		return fmt.Errorf("error while reading client stream: %w", err)
	}
	for _, p := range data {
		ppb := toPortpb(p)
		err = stream.Send(&portpb.PortResponse{Port: &ppb})
		if err != nil {
			s.Log.Errorf("error while reading client stream loop: %v", err)
			return fmt.Errorf("error while reading client stream loop: %w", err)
		}
	}

	return nil
}

func toPortpb(port entity.Port) portpb.Port {
	id := uint32(port.ID)
	return portpb.Port{
		Id:        id,
		Name:      port.Name,
		City:      port.City,
		Country:   port.Country,
		Alias:     port.Alias,
		Regions:   port.Regions,
		Latitude:  port.Latitude,
		Longitude: port.Longitude,
		Province:  port.Province,
		Timezone:  port.Timezone,
		Unlocs:    port.Unlocs,
		Code:      port.Code,
	}
}
