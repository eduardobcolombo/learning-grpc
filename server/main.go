package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/eduardobcolombo/learning-grpc/portpb"
	"google.golang.org/grpc"
)

type server struct {
	portpb.UnimplementedPortServiceServer
}

func main() {
	fmt.Println("Server is running...")

	//	TODO:
	// dbdriver := os.Getenv("DB_DRIVER")
	// host := os.Getenv("DB_HOST")
	// password := os.Getenv("DB_PASSWORD")
	// user := os.Getenv("DB_USER")
	// dbname := os.Getenv("DB_NAME")
	// dbport := os.Getenv("DB_PORT")

	// services, err := persistence.NewRepositories(dbdriver, user, password, dbport, host, dbname)
	// if err != nil {
	// 	panic(err)
	// }
	// defer services.Close()
	// services.Automigrate()

	// ports := interfaces.NewPort(services.Port)

	port := os.Getenv("PORT")
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	portpb.RegisterPortServiceServer(s, &server{})

	if err := s.Serve(list); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (*server) PortsUpdate(stream portpb.PortService_PortsUpdateServer) error {

	fmt.Printf("PortsUpdate function was invoked with a streaming request")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&portpb.PortResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		city := req.GetPort().GetCity()
		result += "Hello " + city

	}

}
