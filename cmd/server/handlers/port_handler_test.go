package handlers

import (
	"context"
	"errors"
	"testing"

	"github.com/eduardobcolombo/learning-grpc/cmd/server/core/port"
	"github.com/eduardobcolombo/learning-grpc/cmd/server/core/port/memory"
	"github.com/eduardobcolombo/learning-grpc/portpb"
)

var (
	errNotFound = errors.New("port not found")
)

func TestPortHandler(t *testing.T) {

	handler := NewPort(port.NewCore(memory.NewStore()))
	ctx := context.Background()

	t.Run("should create a port", func(t *testing.T) {
		portReq := portpb.PortRequest{
			Port: &portpb.Port{
				Name:      "Bordeaux",
				City:      "Bordeaux",
				Country:   "France",
				Alias:     "",
				Regions:   "",
				Latitude:  -0.57918,
				Longitude: 44.837789,
				Province:  "Gironde",
				Timezone:  "Europe/Paris",
				Unlocs:    "FRBOD",
				Code:      "42707",
			},
		}
		if err := handler.Create(ctx, &portReq); err != nil {
			t.Errorf("Error on create port: %v", err)
		}

		port, err := handler.GetByUnloc(ctx, portReq.Port.GetUnlocs())
		if err != nil {
			t.Errorf("Error on get port by unloc: %v", err)
		}

		if port.Name != portReq.Port.GetName() {
			t.Errorf("Expected name %s, got %s", portReq.Port.GetName(), port.Name)
		}
	})

	t.Run("should update a port", func(t *testing.T) {
		portReq := portpb.PortRequest{
			Port: &portpb.Port{
				Name:      "Brest",
				City:      "Brest",
				Country:   "France",
				Alias:     "",
				Regions:   "",
				Latitude:  -4.48,
				Longitude: 48.4,
				Province:  "Finistère",
				Timezone:  "Europe/Paris",
				Unlocs:    "FRBES",
				Code:      "42709",
			},
		}
		if err := handler.Create(ctx, &portReq); err != nil {
			t.Errorf("Error on create port: %v", err)
		}

		port, err := handler.GetByUnloc(ctx, portReq.Port.GetUnlocs())
		if err != nil {
			t.Errorf("Error on get port by unloc: %v", err)
		}

		portReqUpdate := portpb.PortRequest{
			Port: &portpb.Port{
				Name:      "Brest.",
				City:      "Brest.",
				Country:   "France.",
				Alias:     "",
				Regions:   "",
				Latitude:  -4.48,
				Longitude: 48.4,
				Province:  "Finistère.",
				Timezone:  "Europe/Paris.",
				Unlocs:    "FRBES",
				Code:      "42709.",
				Id:        uint32(port.ID),
			},
		}
		if err := handler.Update(ctx, &portReqUpdate); err != nil {
			t.Errorf("Error on update port: %v", err)
		}

		portUpdated, err := handler.GetByUnloc(ctx, portReqUpdate.Port.GetUnlocs())
		if err != nil {
			t.Errorf("Error on get port by unloc: %v", err)
		}

		if portUpdated.Name != portReqUpdate.Port.GetName() {
			t.Errorf("Expected name %s, got %s", portReqUpdate.Port.GetName(), portUpdated.Name)
		}
	})

	t.Run("should get a port by id", func(t *testing.T) {
		portReq := portpb.PortRequest{
			Port: &portpb.Port{
				Name:      "Calais",
				City:      "Calais",
				Country:   "France",
				Alias:     "",
				Regions:   "",
				Latitude:  1.858686,
				Longitude: 50.95137,
				Province:  "Pas-de-Calais",
				Timezone:  "Europe/Paris",
				Unlocs:    "FRCQF",
				Code:      "42713",
			},
		}
		if err := handler.Create(ctx, &portReq); err != nil {
			t.Errorf("Error on create port: %v", err)
		}

		port, err := handler.GetByUnloc(ctx, portReq.Port.GetUnlocs())
		if err != nil {
			t.Errorf("Error on get port by unloc: %v", err)
		}

		portByID, err := handler.GetByID(ctx, uint(port.ID))
		if err != nil {
			t.Errorf("Error on get port by id: %v", err)
		}

		if portByID.Name != port.Name {
			t.Errorf("Expected name %s, got %s", port.Name, portByID.Name)
		}
	})

	t.Run("should get all ports", func(t *testing.T) {
		portReq := portpb.PortRequest{
			Port: &portpb.Port{
				Name:      "Dunkerque",
				City:      "Dunkerque",
				Country:   "France",
				Alias:     "",
				Regions:   "",
				Latitude:  2.377432,
				Longitude: 51.03431,
				Province:  "Nord",
				Timezone:  "Europe/Paris",
				Unlocs:    "FRDKK",
				Code:      "42714",
			},
		}
		if err := handler.Create(ctx, &portReq); err != nil {
			t.Errorf("Error on create port: %v", err)
		}

		ports, err := handler.GetAll(ctx)
		if err != nil {
			t.Errorf("Error on get all ports: %v", err)
		}

		if len(ports) < 1 {
			t.Errorf("Expected at least one port, got none")
		}
	})
	t.Run("should delete a port", func(t *testing.T) {
		portReq := portpb.PortRequest{
			Port: &portpb.Port{
				Name:      "Dunkerque",
				City:      "Dunkerque",
				Country:   "France",
				Alias:     "",
				Regions:   "",
				Latitude:  2.377432,
				Longitude: 51.03431,
				Province:  "Nord",
				Timezone:  "Europe/Paris",
				Unlocs:    "FRDKK",
				Code:      "42714",
			},
		}
		if err := handler.Create(ctx, &portReq); err != nil {
			t.Errorf("Error on create port: %v", err)
		}

		port, err := handler.GetByUnloc(ctx, portReq.Port.GetUnlocs())
		if err != nil {
			t.Errorf("Error on get port by unloc: %v", err)
		}

		portIdReq := &portpb.PortIdRequest{
			Id: uint32(port.ID),
		}

		if err := handler.Delete(ctx, portIdReq); err != nil {
			t.Errorf("Error on delete port: %v", err)
		}

		portByID, err := handler.GetByID(ctx, uint(port.ID))
		if err.Error() != errNotFound.Error() {
			t.Errorf("Expected: %+v, but got %+v", errNotFound, err)
		}

		if portByID.Name != "" {
			t.Errorf("Expected name %s, got %s", "", portByID.Name)
		}

	})

}
