package core_port

import (
	"encoding/json"
	"fmt"
	"io"
	"testing"

	"github.com/eduardobcolombo/learning-grpc/foundation"
	"github.com/eduardobcolombo/learning-grpc/portpb"
	mock_portpb "github.com/eduardobcolombo/learning-grpc/portpb/mock_port"
	"github.com/golang/mock/gomock"
)

func TestToPortRequest(t *testing.T) {
	log, err := foundation.NewLogger(&foundation.LoggerConfig{})
	if err != nil {
		t.Fatalf("error creating logger: %v", err)
	}

	tests := []struct {
		name    string
		input   interface{}
		want    *portpb.PortRequest
		wantErr bool
	}{
		{
			name: "valid input",
			input: map[string]interface{}{
				"name":        "Test Port",
				"city":        "Test City",
				"country":     "Test Country",
				"alias":       []string{"Alias1"},
				"regions":     []string{"Region1"},
				"coordinates": []float64{1.23, 4.56},
				"province":    "Test Province",
				"timezone":    "Test Timezone",
				"unlocs":      []string{"UNLOC1"},
				"code":        "Code1",
			},
			want: &portpb.PortRequest{
				Port: &portpb.Port{
					Name:      "Test Port",
					City:      "Test City",
					Country:   "Test Country",
					Alias:     "Alias1",
					Regions:   "Region1",
					Latitude:  1.23,
					Longitude: 4.56,
					Province:  "Test Province",
					Timezone:  "Test Timezone",
					Unlocs:    "UNLOC1",
					Code:      "Code1",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid input",
			input: map[string]interface{}{
				"name":        "Test Port",
				"city":        "Test City",
				"country":     "Test Country",
				"alias":       "Alias1",     // alias should be a slice
				"regions":     "Region1",    // regions should be a slice
				"coordinates": "1.23, 4.56", // coordinates should be a slice of floats
				"province":    "Test Province",
				"timezone":    "Test Timezone",
				"unlocs":      "UNLOC1", // unlocs should be a slice
				"code":        "Code1",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid input",
			input:   make(chan int),
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toPortRequest(log, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("toPortRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !jsonEqual(got, tt.want) {
				t.Errorf("toPortRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRetrieve(t *testing.T) {
	log, err := foundation.NewLogger(&foundation.LoggerConfig{})
	if err != nil {
		t.Fatalf("error creating logger: %v", err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_portpb.NewMockPortServiceClient(ctrl)
	core := NewCore(log, mockClient)

	tests := []struct {
		name    string
		setup   func()
		want    []*portpb.Port
		wantErr bool
	}{
		{
			name: "successful retrieval",
			setup: func() {
				stream := mock_portpb.NewMockPortService_GetAllClient(ctrl)
				mockClient.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(stream, nil)
				stream.EXPECT().Recv().Return(&portpb.PortResponse{
					Port: &portpb.Port{
						Name:   "Port1",
						Unlocs: "Unloc1",
					},
				}, nil)
				stream.EXPECT().Recv().Return(nil, io.EOF)
			},
			want: []*portpb.Port{
				{
					Name:   "Port1",
					Unlocs: "Unloc1",
				},
			},
			wantErr: false,
		},
		{
			name: "error during GetAll call",
			setup: func() {
				mockClient.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("GetAll error"))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error during stream receive",
			setup: func() {
				stream := mock_portpb.NewMockPortService_GetAllClient(ctrl)
				mockClient.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(stream, nil)
				stream.EXPECT().Recv().Return(nil, fmt.Errorf("stream receive error"))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			got, err := core.Retrieve()
			if (err != nil) != tt.wantErr {
				t.Errorf("Retrieve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !jsonEqual(got, tt.want) {
				t.Errorf("Retrieve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetByUnloc(t *testing.T) {
	log, err := foundation.NewLogger(&foundation.LoggerConfig{})
	if err != nil {
		t.Fatalf("error creating logger: %v", err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_portpb.NewMockPortServiceClient(ctrl)
	core := NewCore(log, mockClient)

	tests := []struct {
		name    string
		unloc   string
		setup   func()
		want    *portpb.Port
		wantErr bool
	}{
		{
			name:  "successful retrieval",
			unloc: "UNLOC1",
			setup: func() {
				mockClient.EXPECT().GetByUnloc(gomock.Any(), &portpb.PortUnlocRequest{Unloc: "UNLOC1"}).Return(&portpb.PortResponse{
					Port: &portpb.Port{
						Name: "Port1",
					},
				}, nil)
			},
			want: &portpb.Port{
				Name: "Port1",
			},
			wantErr: false,
		},
		{
			name:  "error during GetByUnloc call",
			unloc: "UNLOC2",
			setup: func() {
				mockClient.EXPECT().GetByUnloc(gomock.Any(), &portpb.PortUnlocRequest{Unloc: "UNLOC2"}).Return(nil, fmt.Errorf("GetByUnloc error"))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			got, err := core.GetByUnloc(tt.unloc)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByUnloc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !jsonEqual(got, tt.want) {
				t.Errorf("GetByUnloc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func jsonEqual(a, b interface{}) bool {
	aj, _ := json.Marshal(a)
	bj, _ := json.Marshal(b)
	return string(aj) == string(bj)
}
