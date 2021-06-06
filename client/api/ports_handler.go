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
		log.Fatalf("Error while updating ports on server: [%v]", err)
		e.Response(w, http.StatusInternalServerError, err)
		return
	}
	e.Response(w, http.StatusOK, ports)
}

func (e *Environment) UpdatePorts(w http.ResponseWriter, r *http.Request) {
	// readjsonFile()
	msg, err := e.updatePortsOnServer()
	if err != nil {
		log.Fatalf("Error while updating ports on server: [%v]", err)
		e.Response(w, http.StatusInternalServerError, err)
		return
	}

	e.Response(w, http.StatusOK, msg)
}

func (e *Environment) retrievePortsFromServer() (ports []*portpb.Port, err error) {
	req := &portpb.ListPortsRequest{}
	stream, err := e.psc.PortsList(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling retrievePortsFromServer: %v", err)
		return ports, err
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving data from server: %v", err)
		}
		ports = append(ports, res.GetPort())
	}
	return ports, nil
}

func (e *Environment) readjsonFile(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", fileName, err.Error())
	}
	defer f.Close()

	r := bufio.NewReader(f)
	dec := json.NewDecoder(r)

	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	for dec.More() {
		var m map[string]interface{}

		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}
		// Sending over GRPC
		// use e.psc
		fmt.Printf("%v\n", m)
	}

	// read closing bracket
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (e *Environment) updatePortsOnServer() (string, error) {
	coords := &portpb.Coordinates{Lat: 111.111, Long: 222.222}
	var unlocs *portpb.Unlocs
	var unloc []string
	unloc = append(unloc, "UNLOCKKKK")
	unlocs = &portpb.Unlocs{
		Unloc: unloc,
	}

	reqs := []*portpb.PortRequest{
		&portpb.PortRequest{
			Port: &portpb.Port{
				Name:        "XPTO",
				City:        "Canela",
				Country:     "Brazil",
				Alias:       []string{},
				Regions:     []string{},
				Coordinates: coords,
				Province:    "RS",
				Timezone:    "TZ",
				Unlocs:      unlocs,
				Code:        "95680-000",
			},
		},
		&portpb.PortRequest{
			Port: &portpb.Port{
				Name:        "XPTO2",
				City:        "Canela",
				Country:     "Brazil",
				Alias:       []string{},
				Regions:     []string{},
				Coordinates: coords,
				Province:    "RS",
				Timezone:    "TZ",
				Unlocs:      unlocs,
				Code:        "95680-000",
			},
		},
	}

	stream, err := e.psc.PortsUpdate(context.Background())
	if err != nil {
		log.Fatalf("error while calling PortsUpdate: %v", err)
		return "", err
	}
	for _, req := range reqs {
		stream.Send(req)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error receiving response from PortsUpdate: %v", err)
		return "", err
	}
	return res.GetResult(), nil
}
