package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"testing"

	"github.com/eduardobcolombo/learning-grpc/cmd/server/app"
	"github.com/eduardobcolombo/learning-grpc/cmd/server/core/port"
	"github.com/eduardobcolombo/learning-grpc/cmd/server/core/port/db"
	"github.com/eduardobcolombo/learning-grpc/cmd/server/domain/entity"
	"github.com/eduardobcolombo/learning-grpc/cmd/server/handlers"
	"github.com/eduardobcolombo/learning-grpc/foundation"
	"github.com/eduardobcolombo/learning-grpc/foundation/sqlDB"
	"github.com/eduardobcolombo/learning-grpc/portpb"
	"github.com/eduardobcolombo/learning-grpc/test/testhelpers"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/test/bufconn"
)

func dialer(t *testing.T) func(context.Context, string) (net.Conn, error) {
	list := bufconn.Listen(1024 * 1024)

	pgCtx, pgContainer, dbCfg := testhelpers.DatabaseContainer()

	pgCfg := &sqlDB.DBConfig{
		Host:     dbCfg.Host,
		Port:     dbCfg.Port,
		User:     dbCfg.User,
		Password: dbCfg.Password,
		Name:     dbCfg.Name,
	}

	cfgLog := foundation.LoggerConfig{
		Level:       "debug",
		ServiceName: "learningGRPC",
	}

	logger, err := foundation.NewLogger(&cfgLog)
	if err != nil {
		fmt.Printf("initialize logger: %v\n", err)
	}
	logger.Z = logger.Z.With(zap.String("env", "TEST"))

	pg, err := sqlDB.New(pgCfg)
	if err != nil {
		log.Panicf("error initializating the DB: %v", err)
	}

	ctx := context.Background()

	if err := pg.AutoMigrate(ctx, &entity.Port{}); err != nil {
		log.Panicf("error running migrations: %v", err)
	}

	srv := app.Server{
		Port: *handlers.NewPort(port.NewCore(db.NewStore(pg.DB))),
		Log:  logger,
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
		pg.Close()
		if err := pgContainer.Terminate(pgCtx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	})
	return func(context.Context, string) (net.Conn, error) {
		return list.Dial()
	}
}

func getClient(t *testing.T) (context.Context, portpb.PortServiceClient) {
	ctx := context.Background()
	cc, err := grpc.NewClient("localhost", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(dialer(t)))
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

	t.Run("should store a record in the DB", func(t *testing.T) {
		request := &portpb.PortRequest{
			Port: &portpb.Port{
				Name:      "Ajman",
				City:      "Ajman",
				Country:   "United Arab Emirates",
				Alias:     "",
				Regions:   "",
				Latitude:  55.5136433,
				Longitude: 25.4052165,
				Province:  "Ajman",
				Timezone:  "Asia/Dubai",
				Unlocs:    "AEAJM",
				Code:      "52000",
			},
		}

		stream, err := client.UpdateAll(ctx)
		if err != nil {
			t.Errorf("got %q want nil", err)
		}
		err = stream.Send(request)
		if err != nil {
			t.Errorf("got %q want nil", err)
		}
		res, err := stream.CloseAndRecv()
		if err != nil {
			t.Errorf("got %q want nil", err)
		}

		got := res.GetResult()
		want := "Received 1 records."
		if got != want {
			t.Errorf("got %s[%T] want %s[%T]", got, got, want, want)
		}

	})

	t.Run("should retrieve the ports list", func(t *testing.T) {

		var ports []*portpb.Port
		req := &portpb.PortsGetAllRequest{}
		stream, err := client.GetAll(ctx, req)
		if err != nil {
			t.Errorf("got %q want nil", err)
		}

		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				t.Errorf("got %q want nil", err)
			}
			ports = append(ports, res.GetPort())
		}

		got := len(ports) > 0
		if got != true {
			fmt.Printf("%v\n", ports)
			t.Errorf("got %t want true", got)
		}
	})

}
