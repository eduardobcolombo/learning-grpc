package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/eduardobcolombo/learning-grpc/server/infrastructure/persistence"
	"github.com/eduardobcolombo/learning-grpc/server/interfaces"
	"github.com/eduardobcolombo/portpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("Server is running...")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	port := "0.0.0.0:" + os.Getenv("PORT")
	fmt.Printf("-----> %s\n", port)
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf("Failed to listen: %v", err)
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
			return
		}
		opts = append(opts, grpc.Creds(crds))
	}
	services, err := initDB()
	if err != nil {
		log.Printf("Error initializating the DB: %v", err)
		return
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
	os.Exit(0)
}

func initDB() (services *persistence.Repositories, err error) {
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")

	services, err = persistence.NewRepositories(user, password, dbport, host, dbname)
	if err != nil {
		return services, err
	}
	services.Automigrate()

	return services, nil
}
