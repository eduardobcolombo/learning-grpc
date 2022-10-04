package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/eduardobcolombo/learning-grpc/cmd/server/domain/entity"
	"github.com/eduardobcolombo/learning-grpc/cmd/server/infrastructure/persistence"
	"github.com/eduardobcolombo/learning-grpc/cmd/server/interfaces"
	"github.com/eduardobcolombo/learning-grpc/internal/pkg/db"
	"github.com/eduardobcolombo/learning-grpc/internal/pkg/foundation"
	"github.com/eduardobcolombo/learning-grpc/internal/pkg/portpb"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func main() {

	// Construct the application logger.
	log, err := foundation.NewLogger("SERVER")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()

	code := run(log)
	os.Exit(code)

}

type Config struct {
	GRPC     GRPC
	DBConfig *db.DBConfig
}

type GRPC struct {
	Host string `envconfig:"GRPC_HOST" default:"0.0.0.0"`
	Port string `envconfig:"GRPC_PORT" default:"50053"`
}

func run(log *zap.SugaredLogger) int {
	log.Infof("Server is running...\n")

	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Fatal(err)
		return 1
	}

	uri := cfg.GRPC.Host + ":" + cfg.GRPC.Port
	log.Infof("-----> %s\n", uri)
	list, err := net.Listen("tcp", uri)
	if err != nil {
		log.Error("Failed to listen: %v", err)
		return 1
	}

	opts := []grpc.ServerOption{}
	// I'm using tls always false, but here is the implementation if we would like to make it tls based
	tls := false
	if tls {
		cFile := "PATH_TO_THE_CERT.crt"
		kFile := "PATH_TO_THE_CERT.pem"
		crds, err := credentials.NewServerTLSFromFile(cFile, kFile)
		if err != nil {
			log.Error("Failed loading certificates: %v", err)
			return 1
		}
		opts = append(opts, grpc.Creds(crds))
	}

	pg, err := db.New(cfg.DBConfig)
	if err != nil {
		log.Error("error initializating the DB: %v", err)
		return 1
	}

	if err := pg.Automigrate(&entity.Port{}, &entity.Alias{}, &entity.Coordinate{}, &entity.Region{}, &entity.Unloc{}); err != nil {
		log.Error("error running migrations: %v", err)
		return 1
	}

	repositories, err := persistence.New(pg)
	if err != nil {
		log.Error("error creating the persistence: %v", err)
		return 1
	}

	srv := interfaces.Server{
		Services: repositories,
	}

	s := grpc.NewServer(opts...)
	portpb.RegisterPortServiceServer(s, &srv)
	reflection.Register(s)

	go func() {
		if err := s.Serve(list); err != nil {
			log.Error("failed to serve: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	log.Info("Stopping the server\n")
	s.Stop()
	log.Info("Closing the listener")
	list.Close()
	log.Info("Closing DB")
	srv.Services.Close()
	log.Info("Shutdown")

	return 0
}
