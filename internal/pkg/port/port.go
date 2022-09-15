package port

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/eduardobcolombo/learning-grpc/internal/pkg/portpb"
	"go.uber.org/zap"
)

type Core struct {
	log *zap.SugaredLogger
	psc portpb.PortServiceClient
}

// NewCore constructs a core for cardtoken api access.
func NewCore(log *zap.SugaredLogger, psc portpb.PortServiceClient) Core {
	return Core{log: log, psc: psc}
}

func (c Core) RetrievePortsFromServer() (ports []*portpb.Port, err error) {
	req := &portpb.ListPortsRequest{}
	stream, err := c.psc.PortsList(context.Background(), req)
	if err != nil {
		c.log.Errorf("error while calling retrievePortsFromServer: %v", err)
		return ports, err
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			c.log.Errorf("Error receiving data from server: %v", err)
			return ports, err
		}
		ports = append(ports, res.GetPort())
	}
	return ports, nil
}

func (c Core) UpdatePortsOnServer(fileName string) (string, error) {

	f, err := os.Open(fileName)
	if err != nil {
		c.log.Errorf("Error to read [file=%v]: %v", fileName, err.Error())
	}
	defer f.Close()
	stream, err := c.psc.PortsUpdate(context.Background())
	if err != nil {
		c.log.Errorf("error while calling PortsUpdate: %v", err)
		return "", err
	}

	// If the JSON file was an JSON array it could work better
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

			jsonbody, err := json.Marshal(v)
			if err != nil {
				// do error check
				fmt.Println(err)
				return "", err
			}

			filledPort, err := fillPortpbWithJSON(c.log, jsonbody)
			if err != nil {
				c.log.Errorf("Error to read [file=%v]: %v", fileName, err.Error())
			}

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

func fillPortpbWithJSON(log *zap.SugaredLogger, jsonbody []byte) (req *portpb.PortRequest, err error) {

	dataPort := portRequest{}
	if err := json.Unmarshal(jsonbody, &dataPort); err != nil {
		// do error check
		log.Errorf("Error reading the JSON file: [%v]", err)
		return req, err
	}

	coord := &portpb.Coordinates{}
	if len(dataPort.Coordinates) == 2 {
		coord = &portpb.Coordinates{Lat: dataPort.Coordinates[0], Long: dataPort.Coordinates[1]}
	}

	return &portpb.PortRequest{
		Port: &portpb.Port{
			Name:        dataPort.Name,
			City:        dataPort.City,
			Country:     dataPort.Country,
			Alias:       dataPort.Alias,
			Regions:     dataPort.Regions,
			Coordinates: coord,
			Province:    dataPort.Province,
			Timezone:    dataPort.Timezone,
			Unlocs: &portpb.Unlocs{
				Unloc: dataPort.Unlocs,
			},
			Code: dataPort.Code,
		},
	}, nil

}

type portRequest struct {
	Name        string    `json:"name"`
	Coordinates []float32 `json:"coordinates"`
	City        string    `json:"city"`
	Province    string    `json:"province"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}
