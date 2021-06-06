package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/eduardobcolombo/portpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Environment struct {
	tls bool
	psc portpb.PortServiceClient
	cc  *grpc.ClientConn
}

func (e *Environment) Response(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte(err.Error())); err != nil {
			log.Fatalf("Error trying to write the error response body: %s ", err)
		}
		return
	}
	// TODO: check the response format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err = w.Write([]byte(response)); err != nil {
		log.Fatalf("Error trying to write the response body: %s ", err)
	}
}

func (e *Environment) GetGRPC() error {
	address := os.Getenv("HOST") + ":" + os.Getenv("PORT")
	fmt.Println("Starting client GRPC connection ", address)

	// I'm using e.tls always false, but here is the implementation if we would like to make it tls based
	opts := grpc.WithInsecure()
	if e.tls {
		cFile := "ADD_THE_CERTIFICATE_PATH_HERE"
		crds, err := credentials.NewClientTLSFromFile(cFile, "")
		if err != nil {
			log.Fatalf("Error loading certificate: %v", err)
			return err
		}
		opts = grpc.WithTransportCredentials(crds)
	}

	cc, err := grpc.Dial(address, opts)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return err
	}
	e.psc = portpb.NewPortServiceClient(cc)
	e.cc = cc

	return nil
}
