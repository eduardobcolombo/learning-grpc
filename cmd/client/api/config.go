package api

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/eduardobcolombo/learning-grpc/internal/pkg/portpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	tls      bool
	psc      portpb.PortServiceClient
	GRPCHost string `envconfig:"GRPC_HOST"`
	GRPCPort string `envconfig:"GRPC_PORT"`
	APIPort  string `envconfig:"API_PORT"`
}

func GRPCInit(cfg *Config) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cfg.GRPCHost = "server"
	address := cfg.GRPCHost + ":" + cfg.GRPCPort
	fmt.Println("Starting client GRPC connection ", address)

	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// I'm using cfg.tls always false, but here is the implementation if we would like to make it tls based
	if cfg.tls {
		cFile := "ADD_THE_CERTIFICATE_PATH_HERE"
		crds, err := credentials.NewClientTLSFromFile(cFile, "")
		if err != nil {
			log.Printf("Error loading certificate: %v", err)
			return nil, err
		}
		opts = append(opts, grpc.WithTransportCredentials(crds))
	}

	cc, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		log.Printf("did not connect: %s %v", address, err)
		return nil, err
	}

	return cc, nil
}
