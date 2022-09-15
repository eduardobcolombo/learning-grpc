package api

import (
	"github.com/eduardobcolombo/learning-grpc/internal/pkg/port"
	"github.com/eduardobcolombo/learning-grpc/internal/pkg/portpb"
)

type Config struct {
	tls      bool
	psc      portpb.PortServiceClient
	GRPCHost string `envconfig:"GRPC_HOST" default:"grpc-server.grpc"`
	GRPCPort string `envconfig:"GRPC_PORT" default:"50053"`
	APIPort  string `envconfig:"API_PORT" default:"8888"`
}

type CoreConfig struct {
	Port port.Core
}
