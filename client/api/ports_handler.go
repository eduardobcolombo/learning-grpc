package api

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/eduardobcolombo/learning-grpc/portpb"
)

func (e *Environment) RetrievePorts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hit here")
	e.GetGRPC()
	doClientStreaming(e.psc)
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

func doClientStreaming(c portpb.PortServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")
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

	stream, err := c.PortsUpdate(context.Background())
	if err != nil {
		log.Fatalf("error while calling PortsUpdate: %v", err)
		return
	}
	for _, req := range reqs {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error receiving response from PortsUpdate: %v", err)
		return
	}
	fmt.Printf("PortsUpdate Response %v\n", res)
}
