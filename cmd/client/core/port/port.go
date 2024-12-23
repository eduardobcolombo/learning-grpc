//go:generate mockgen --destination=./mock_port/mock_port.go --source=./port.go
package core_port

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/eduardobcolombo/learning-grpc/foundation"
	"github.com/eduardobcolombo/learning-grpc/portpb"
)

// CoreService defines the interface for the Core methods.
type CoreService interface {
	Create(port PortCore) (*PortCore, error)
	Retrieve() ([]*portpb.Port, error)
	GetByUnloc(unloc string) (*portpb.Port, error)
	Update(fileName string) (string, error)
}

type Core struct {
	log *foundation.Logger
	psc portpb.PortServiceClient
}

// NewCore constructs a core for cardtoken api access.
func NewCore(log *foundation.Logger, psc portpb.PortServiceClient) Core {
	return Core{log: log, psc: psc}
}

func (c Core) Create(port PortCore) (*PortCore, error) {
	return &port, nil
}

// Retrieve call GRPC server to retrieve all the ports.
func (c Core) Retrieve() (ports []*portpb.Port, err error) {
	req := &portpb.PortsGetAllRequest{}
	stream, err := c.psc.GetAll(context.Background(), req)
	if err != nil {
		err = fmt.Errorf("error while calling Retrieve: %w", err)
		return ports, err
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			err = fmt.Errorf("error receiving data from server: %w", err)
			return ports, err
		}
		ports = append(ports, res.GetPort())
	}

	return ports, nil
}

// Retrieve call GRPC server to retrieve all the ports.
func (c Core) GetByUnloc(unloc string) (port *portpb.Port, err error) {
	req := &portpb.PortUnlocRequest{
		Unloc: unloc,
	}
	res, err := c.psc.GetByUnloc(context.Background(), req)
	if err != nil {
		err = fmt.Errorf("error while calling Retrieve: %w", err)
		return nil, err
	}

	fmt.Printf("%+v\n", res.GetPort())

	return res.GetPort(), nil
}

// Update send the JSON file to the server using GRPC.
func (c Core) Update(fileName string) (string, error) {

	f, err := os.Open(fileName)
	if err != nil {
		err = fmt.Errorf("error to read [file=%v]: %w", fileName, err)
		return "", err
	}
	defer f.Close()
	stream, err := c.psc.UpdateAll(context.Background())
	if err != nil {
		c.log.Errorf("error while calling Update: %v", err)
		return "", err
	}

	// If the JSON file was an JSON array it could work better.
	r := bufio.NewReader(f)
	dec := json.NewDecoder(r)

	for dec.More() {
		var m map[string]interface{}
		err := dec.Decode(&m)
		if err != nil {
			return "", err
		}

		var count int

		for _, v := range m {

			filledPort, err := toPortRequest(c.log, v)
			if err != nil {
				c.log.Errorf("Error to read [file=%v]: %v", fileName, err.Error())
			}

			c.log.Infof("Sending port: %v", filledPort)

			if err := stream.Send(filledPort); err != nil {
				c.log.Errorf("Error to send stream: %s", err.Error())
			}

			count += 1
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		c.log.Errorf("error receiving response from PortsUpdate: %v", err)
		return "", err
	}
	return res.GetResult(), nil
}

// toPortRequest parse the Port with JSON to PB type.
func toPortRequest(log *foundation.Logger, v interface{}) (req *portpb.PortRequest, err error) {

	jsonbody, err := json.Marshal(v)
	if err != nil {
		// do error check.
		fmt.Println(err)
		return nil, err
	}

	jsonPort := PortJsonRequest{}
	if err := json.Unmarshal(jsonbody, &jsonPort); err != nil {
		// do error check.
		log.Errorf("Error reading the JSON file: [%v]", err)
		return req, err
	}
	dataPort := PortCore{}
	dataPort.Name = jsonPort.Name
	if len(jsonPort.Coordinates) > 0 {
		dataPort.Latitude = jsonPort.Coordinates[0]
	}
	if len(jsonPort.Coordinates) > 1 {
		dataPort.Longitude = jsonPort.Coordinates[1]
	}
	dataPort.City = jsonPort.City
	dataPort.Province = jsonPort.Province
	dataPort.Country = jsonPort.Country
	if len(jsonPort.Alias) > 0 {
		dataPort.Alias = jsonPort.Alias[0]
	}
	if len(jsonPort.Regions) > 0 {
		dataPort.Regions = jsonPort.Regions[0]
	}
	dataPort.Timezone = jsonPort.Timezone
	if len(jsonPort.Unlocs) > 0 {
		dataPort.Unlocs = jsonPort.Unlocs[0]
	}
	dataPort.Code = jsonPort.Code

	return &portpb.PortRequest{
		Port: dataPort.toPortpb(),
	}, nil

}

type PortJsonRequest struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

type PortCore struct {
	Name      string
	Latitude  float64
	Longitude float64
	City      string
	Province  string
	Country   string
	Alias     string
	Regions   string
	Timezone  string
	Unlocs    string
	Code      string
}

func (p PortCore) toPortpb() *portpb.Port {

	return &portpb.Port{
		Name:      p.Name,
		City:      p.City,
		Country:   p.Country,
		Alias:     p.Alias,
		Regions:   p.Regions,
		Latitude:  p.Latitude,
		Longitude: p.Longitude,
		Province:  p.Province,
		Timezone:  p.Timezone,
		Unlocs:    p.Unlocs,
		Code:      p.Code,
	}
}
