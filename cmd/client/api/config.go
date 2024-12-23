package api

import (
	core_port "github.com/eduardobcolombo/learning-grpc/cmd/client/core/port"
	"github.com/eduardobcolombo/learning-grpc/foundation"
	"github.com/eduardobcolombo/learning-grpc/portpb"
)

type Config struct {
	TLS         bool
	PSC         portpb.PortServiceClient
	GRPCHost    string `envconfig:"GRPC_HOST" default:"docker.internal"`
	GRPCPort    string `envconfig:"GRPC_PORT" default:"50053"`
	APIPort     string `envconfig:"API_PORT" default:"8888"`
	Log         foundation.LoggerConfig
	Environment string `envconfig:"ENVIRONMENT" default:"production"`
}

type CoreConfig struct {
	Port core_port.CoreService
}

func NewCoreConfig(log *foundation.Logger, psc portpb.PortServiceClient) CoreConfig {
	return CoreConfig{Port: core_port.NewCore(log, psc)}
}
