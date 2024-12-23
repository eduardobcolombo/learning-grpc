package port

import (
	"context"
	"errors"
	"testing"

	"github.com/eduardobcolombo/learning-grpc/cmd/server/core/port/memory"
	"github.com/eduardobcolombo/learning-grpc/cmd/server/domain/entity"
)

var (
	errNotFound = errors.New("port not found")
)

func TestCorePort(t *testing.T) {

	core := NewCore(memory.NewStore())
	t.Run("should create a port", func(t *testing.T) {
		if err := core.Create(context.Background(), entity.Port{Name: "Guaibas Port", Unlocs: "BRPOA"}); err != nil {
			t.Errorf("Error on Create: %v", err)
		}

		port, err := core.GetByUnloc(context.Background(), "BRPOA")
		if err != nil {
			t.Errorf("Error on GetByUnloc: %v", err)
		}
		if port.Name != "Guaibas Port" {
			t.Errorf("Expected 'Guaibas Port' got %v", port.Name)
		}
		if port.Unlocs != "BRPOA" {
			t.Errorf("Expected 'BRPOA' got %v", port.Unlocs)
		}

	})

	t.Run("should update a port", func(t *testing.T) {
		if err := core.Create(context.Background(), entity.Port{Name: "Guaibas Port", Unlocs: "BRPOA"}); err != nil {
			t.Errorf("Error on Create: %v", err)
		}

		if err := core.Update(context.Background(), entity.Port{ID: 1, Name: "Guaibas Port UPDATED", Unlocs: "BRPOA"}); err != nil {
			t.Errorf("Error on Update: %v", err)
		}

		port, err := core.GetByUnloc(context.Background(), "BRPOA")
		if err != nil {
			t.Errorf("Error on GetByUnloc: %v", err)
		}
		if port.Name != "Guaibas Port UPDATED" {
			t.Errorf("Expected 'Guaibas Port UPDATED' got %v", port.Name)
		}
		if port.Unlocs != "BRPOA" {
			t.Errorf("Expected 'BRPOA' got %v", port.Unlocs)
		}
	})

	t.Run("should delete a port", func(t *testing.T) {
		if err := core.Create(context.Background(), entity.Port{Name: "Guaibas Port", Unlocs: "BRPOA"}); err != nil {
			t.Errorf("Error on Create: %v", err)
		}

		portUpdated, err := core.GetByUnloc(context.Background(), "BRPOA")
		if err == errNotFound {
			t.Errorf("Expected error got nil")
		}

		if err := core.Delete(context.Background(), portUpdated.ID); err != nil {
			t.Errorf("Error on Delete: %v", err)
		}

		if _, err := core.GetByUnloc(context.Background(), "BRPOA"); err == errNotFound {
			t.Errorf("Expected error got nil")
		}
	})

	t.Run("should get all ports", func(t *testing.T) {
		if err := core.Create(context.Background(), entity.Port{Name: "Guaibas Port", Unlocs: "BRPOA"}); err != nil {
			t.Errorf("Error on Create: %v", err)
		}

		ports, err := core.GetAll(context.Background())
		if err != nil {
			t.Errorf("Error on GetAll: %v", err)
		}
		if len(ports) < 1 {
			t.Errorf("Expected 1 got %v", len(ports))
		}
	})

	t.Run("should get a port by ID", func(t *testing.T) {
		if err := core.Create(context.Background(), entity.Port{Name: "Guaibas Port", Unlocs: "BRPOA"}); err != nil {
			t.Errorf("Error on Create: %v", err)
		}

		portGet, err := core.GetByUnloc(context.Background(), "BRPOA")
		if err == errNotFound {
			t.Errorf("Expected error got nil")
		}

		port, err := core.GetByID(context.Background(), portGet.ID)
		if err != nil {
			t.Errorf("Error on GetByID: %v", err)
		}
		if port.Name != "Guaibas Port" {
			t.Errorf("Expected 'Guaibas Port' got %v", port.Name)
		}
		if port.Unlocs != "BRPOA" {
			t.Errorf("Expected 'BRPOA' got %v", port.Unlocs)
		}
	})

}
