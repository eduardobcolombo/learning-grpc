package api

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/eduardobcolombo/portpb"
)

func (e *Environment) RetrievePorts(w http.ResponseWriter, r *http.Request) {
	ports, err := e.retrievePortsFromServer()
	if err != nil {
		log.Printf("Error while updating ports on server: [%v]", err)
		e.Response(w, http.StatusInternalServerError, err)
		return
	}
	e.Response(w, http.StatusOK, ports)
}

func (e *Environment) UpdatePorts(w http.ResponseWriter, r *http.Request) {
	// TODO: Read the file from upload or URL if it is the case
	fileName := "../ports.json"
	fmt.Printf("Importing the file: %s\n\n", fileName)
	msg, err := e.UpdatePortsOnServer(fileName)
	if err != nil {
		log.Printf("Error while updating ports on server: [%v]", err)
		e.Response(w, http.StatusInternalServerError, err)
		return
	}

	e.Response(w, http.StatusOK, msg)
}

func (e *Environment) retrievePortsFromServer() (ports []*portpb.Port, err error) {
	req := &portpb.ListPortsRequest{}
	stream, err := e.psc.PortsList(context.Background(), req)
	if err != nil {
		log.Printf("error while calling retrievePortsFromServer: %v", err)
		return ports, err
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error receiving data from server: %v", err)
			return ports, err
		}
		ports = append(ports, res.GetPort())
	}
	return ports, nil
}

func (e *Environment) UpdatePortsOnServer(fileName string) (string, error) {

	f, err := os.Open(fileName)
	if err != nil {
		log.Printf("Error to read [file=%v]: %v", fileName, err.Error())
	}
	defer f.Close()
	stream, err := e.psc.PortsUpdate(context.Background())
	if err != nil {
		log.Printf("error while calling PortsUpdate: %v", err)
		return "", err
	}

	// If the JSON file was an JSON array it could work better
	r := bufio.NewReader(f)
	dec := json.NewDecoder(r)

	for dec.More() {
		var m map[string]interface{}
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range m {

			jsonbody, err := json.Marshal(v)
			if err != nil {
				// do error check
				fmt.Println(err)
				return "", err
			}

			filledPort, err := fillPortpbWithJSON(jsonbody)
			if err != nil {
				log.Printf("Error to read [file=%v]: %v", fileName, err.Error())
			}

			stream.Send(filledPort)

		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("error receiving response from PortsUpdate: %v", err)
		return "", err
	}
	return res.GetResult(), nil
}

func fillPortpbWithJSON(jsonbody []byte) (req *portpb.PortRequest, err error) {

	dataPort := portRequest{}
	if err := json.Unmarshal(jsonbody, &dataPort); err != nil {
		// do error check
		log.Printf("Error reading the JSON file: [%v]", err)
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
