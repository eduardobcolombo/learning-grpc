package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/eduardobcolombo/learning-grpc/cmd/server/infrastructure/persistence"
	"github.com/eduardobcolombo/learning-grpc/cmd/server/interfaces"
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
	GRPC GRPC
	DB   persistence.DB
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

	services, err := InitDB(&cfg.DB)
	if err != nil {
		log.Error("Error initializating the DB: %v", err)
		return 1
	}

	srv := interfaces.Server{
		Services: services,
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
	services.Close()
	log.Info("Shutdown")

	return 0
}

func InitDB(cfg *persistence.DB) (services *persistence.Repositories, err error) {
	host := cfg.Host
	password := cfg.Password
	user := cfg.User
	dbname := cfg.Name
	dbport := cfg.Port

	services, err = persistence.NewRepositories(user, password, dbport, host, dbname)
	if err != nil {
		return services, err
	}

	if err := services.Automigrate(); err != nil {
		return services, err
	}

	return services, nil
}
