package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/eduardobcolombo/learning-grpc/cmd/server/infrastructure/persistence"
	"github.com/eduardobcolombo/learning-grpc/cmd/server/interfaces"
	"github.com/eduardobcolombo/learning-grpc/internal/pkg/portpb"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func main() {

	code := run()
	os.Exit(code)

}

type Config struct {
	GRPC GRPC
	DB   DB
}

type GRPC struct {
	Host string `envconfig:"GRPC_HOST"`
	Port string `envconfig:"GRPC_PORT"`
}
type DB struct {
	Host     string `envconfig:"DB_HOST"`
	Password string `envconfig:"DB_PASSWORD"`
	User     string `envconfig:"DB_USER"`
	Name     string `envconfig:"DB_NAME"`
	Port     string `envconfig:"DB_PORT"`
}

func run() int {
	fmt.Println("Server is running...")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Fatal(err)
		return 1
	}

	uri := cfg.GRPC.Host + ":" + cfg.GRPC.Port
	fmt.Printf("-----> %s\n", uri)
	list, err := net.Listen("tcp", uri)
	if err != nil {
		log.Printf("Failed to listen: %v", err)
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
			log.Printf("Failed loading certificates: %v", err)
			return 1
		}
		opts = append(opts, grpc.Creds(crds))
	}
	services, err := InitDB(&cfg.DB)
	if err != nil {
		log.Printf("Error initializating the DB: %v", err)
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
			log.Printf("failed to serve: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Closing the listener")
	list.Close()
	fmt.Println("Closing DB")
	services.Close()
	fmt.Println("Shutdown")

	return 0
}

func InitDB(cfg *DB) (services *persistence.Repositories, err error) {
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
