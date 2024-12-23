package memory

import (
	"context"
	"testing"

	"github.com/eduardobcolombo/learning-grpc/cmd/server/domain/entity"
)

// These tests uses native testing pkg to show up how to use it

func TestCreateAndUpdatePort(t *testing.T) {
	s := NewStore()

	tests := []struct {
		title        string
		originalPort entity.Port
		updatePort   entity.Port
		expectedPort entity.Port
	}{
		{
			title: "should create and updated a port",
			originalPort: entity.Port{
				Name:      "Guaibas Port",
				City:      "Porto Alegre",
				Country:   "Brazil",
				Alias:     "GPOA",
				Regions:   "South",
				Latitude:  1.12345,
				Longitude: 5.43211,
				Province:  "Porto Alegre",
				Timezone:  "",
				Unlocs:    "BRPOA",
				Code:      "12345",
			},
			updatePort: entity.Port{
				Name:   "Guaibas Port UPDATED",
				Unlocs: "BRPOA",
			},
			expectedPort: entity.Port{
				Name:   "Guaibas Port UPDATED",
				Unlocs: "BRPOA",
			},
		},
		{
			title: "should create and updated a port",
			originalPort: entity.Port{
				Name:      "Rio de Janeiro Port",
				City:      "Rio de Janeiro",
				Country:   "Brazil",
				Alias:     "RJ",
				Regions:   "South",
				Latitude:  1.22345,
				Longitude: 5.33211,
				Province:  "Rio de Janeiro",
				Timezone:  "",
				Unlocs:    "BRRJB",
				Code:      "2342",
			},
			updatePort: entity.Port{
				Name:   "Rio de Janeiro Port UPDATED",
				Unlocs: "BRRJB",
			},
			expectedPort: entity.Port{
				Name:   "Rio de Janeiro Port UPDATED",
				Unlocs: "BRRJB",
			},
		},
	}

	ctx := context.Background()

	for _, test := range tests {
		tt := test
		t.Run(tt.title, func(t *testing.T) {
			if err := s.Create(ctx, tt.originalPort); err != nil {
				t.Error("expected create a new port")
			}

			if _, err := s.GetByUnloc(ctx, "INVALID_UNLOC"); err != entity.ErrPortNotFound {
				t.Errorf("expected %s, but got %v", entity.ErrPortNotFound, err)
			}

			portPOA, err := s.GetByUnloc(ctx, tt.originalPort.Unlocs)
			if err != nil {
				t.Errorf("expected err nil, but got %v", err)
			}

			if portPOA.ID < 1 {
				t.Error("expected port.ID > than 0")
			}

			tt.updatePort.ID = portPOA.ID
			if err := s.Update(ctx, tt.updatePort); err != nil {
				t.Errorf("expected to update ports, but got %v", err)
			}

			if err := s.Update(ctx, entity.Port{ID: 999}); err != entity.ErrPortNotFound {
				t.Errorf("expected %s, but got %v", entity.ErrPortNotFound, err)
			}

			if _, err := s.GetByID(ctx, 999); err != entity.ErrPortNotFound {
				t.Errorf("expected %s, but got %v", entity.ErrPortNotFound, err)
			}

			updatedPort, err := s.GetByID(ctx, portPOA.ID)
			if err != nil {
				t.Errorf("expected err nil, but got %v", err)
			}

			if updatedPort.Name != tt.expectedPort.Name {
				t.Errorf("%d expected %s, but got %s", portPOA.ID, tt.expectedPort.Name, updatedPort.Name)
			}

			portData, err := s.GetByUnloc(ctx, tt.updatePort.Unlocs)
			if err != nil {
				t.Errorf("expected err nil, but got %v", err)
			}

			var portUpdate entity.Port
			portUpdate.ID = portData.ID
			portUpdate.Name = "Eduardos Port"
			if err := s.Update(ctx, portUpdate); err != nil {
				t.Errorf("expected to update ports, but got %v", err)
			}

		})
	}

	t.Run("should return err port not found", func(t *testing.T) {
		if _, err := s.GetByID(ctx, 999); err != entity.ErrPortNotFound {
			t.Errorf("expected %s, but got %v", entity.ErrPortNotFound, err)
		}
	})

	t.Run("should return all ports", func(t *testing.T) {
		allPorts, err := s.GetAll(ctx)
		if err != nil {
			t.Errorf("expected err nil, but got %v", err)
		}

		if len(tests) != len(allPorts) {
			t.Error("expected all ports, but get nothing")
		}
	})

	t.Run("should delete a port", func(t *testing.T) {
		allPorts, err := s.GetAll(ctx)
		if err != nil {
			t.Errorf("expected err nil, but got %v", err)
		}

		for _, p := range allPorts {
			port := p
			if err := s.Delete(ctx, 999); err != entity.ErrPortNotFound {
				t.Errorf("expected %s, but got %v", entity.ErrPortNotFound, err)
			}

			if err := s.Delete(ctx, port.ID); err != nil {
				t.Errorf("expected err nil, but got %v", err)
			}

			deletedPort, _ := s.GetByID(ctx, port.ID)
			if deletedPort.ID > 0 {
				t.Errorf("expected nothing, but got %v", deletedPort)
			}
		}

	})
}
