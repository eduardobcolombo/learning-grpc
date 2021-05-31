package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/academiadaweb/learning-grpc/portpb"
	"google.golang.org/grpc"
)

func main() {
	address := os.Getenv("HOST") + os.Getenv("PORT")
	fmt.Println("Client -> ", address)

	cc, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer cc.Close()
	c := portpb.NewPortServiceClient(cc)
	doClientStreaming(c)

	// read the JSON file
	readjsonFile("../ports.json")

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
	}
	for _, req := range reqs {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error receiving response from PortsUpdate: %v", err)
	}
	fmt.Printf("PortsUpdate Response %v\n", res)
}

func readjsonFile(fileName string) error {

	start := time.Now()
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", fileName, err.Error())
	}

	r := bufio.NewReader(f)
	dec := json.NewDecoder(r)

	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

	for dec.More() {
		var m map[string]interface{}

		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", m)
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

	elapsed := time.Since(start)

	fmt.Printf("To parse the file took [%v]\n", elapsed)
	return nil
}
