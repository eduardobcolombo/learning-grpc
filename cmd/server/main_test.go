package main

import (
	"context"
	"io"
	"log"
	"net"
	"testing"

	"github.com/eduardobcolombo/learning-grpc/cmd/server/infrastructure/persistence"
	"github.com/eduardobcolombo/learning-grpc/cmd/server/interfaces"
	"github.com/eduardobcolombo/learning-grpc/internal/pkg/db"
	"github.com/eduardobcolombo/learning-grpc/internal/pkg/portpb"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/test/bufconn"
)

var assertCorrectMessage = func(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %s[%T] want %s[%T]", got, got, want, want)
	}
}
var assertNil = func(t *testing.T, got interface{}) {
	t.Helper()
	if got != nil {
		t.Errorf("got %q want nil", got)
	}
}
var assertTrue = func(t *testing.T, got interface{}) {
	t.Helper()
	if got != true {
		t.Errorf("got %q want true", got)
	}
}

func dialer(t *testing.T) func(context.Context, string) (net.Conn, error) {
	list := bufconn.Listen(1024 * 1024)

	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Fatal(err)
	}

	pg, err := db.New(cfg.DBConfig)
	if err != nil {
		log.Panicf("Error initializating the DB: %v", err)
	}

	repositories, err := persistence.New(pg)
	if err != nil {
		log.Panicf("error creating the persistence: %v", err)
	}

	srv := interfaces.Server{
		Services: repositories,
	}

	s := grpc.NewServer()
	portpb.RegisterPortServiceServer(s, &srv)
	reflection.Register(s)

	go func() {
		if err := s.Serve(list); err != nil {
			log.Fatal(err)
		}
	}()
	t.Cleanup(func() {
		s.Stop()
		list.Close()
		srv.Services.Close()
	})
	return func(context.Context, string) (net.Conn, error) {
		return list.Dial()
	}
}

func getClient(t *testing.T) (context.Context, portpb.PortServiceClient) {
	ctx := context.Background()
	cc, err := grpc.DialContext(ctx, "", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(dialer(t)))
	if err != nil {
		log.Fatal(err)
	}
	t.Cleanup(func() {
		defer cc.Close()
	})
	client := portpb.NewPortServiceClient(cc)
	return ctx, client
}
func TestPortsUpdate(t *testing.T) {
	ctx, client := getClient(t)

	t.Run("Storing a record in the DB", func(t *testing.T) {
		request := &portpb.PortRequest{
			Port: &portpb.Port{
				Name:        "Ajman",
				City:        "Ajman",
				Country:     "United Arab Emirates",
				Alias:       []string{},
				Regions:     []string{},
				Coordinates: &portpb.Coordinates{Lat: 55.5136433, Long: 25.4052165},
				Province:    "Ajman",
				Timezone:    "Asia/Dubai",
				Unlocs:      &portpb.Unlocs{Unloc: []string{"AEAJM"}},
				Code:        "52000",
			},
		}

		stream, err := client.PortsUpdate(ctx)
		assertNil(t, err)
		err = stream.Send(request)
		assertNil(t, err)
		res, err := stream.CloseAndRecv()
		assertNil(t, err)
		assertCorrectMessage(t, res.GetResult(), "Received 1 records.")

	})

}

func TestPortsList(t *testing.T) {
	ctx, client := getClient(t)

	t.Run("Retrieving the ports list", func(t *testing.T) {

		var ports []*portpb.Port
		req := &portpb.ListPortsRequest{}
		stream, err := client.PortsList(ctx, req)
		assertNil(t, err)
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			assertNil(t, err)
			ports = append(ports, res.GetPort())
		}

		assertTrue(t, len(ports) > 0)
	})

}
