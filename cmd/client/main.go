package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/eduardobcolombo/learning-grpc/cmd/client/api"
	"github.com/eduardobcolombo/learning-grpc/cmd/client/core/port"
	"github.com/eduardobcolombo/learning-grpc/internal/pkg/foundation"
	"github.com/eduardobcolombo/learning-grpc/internal/pkg/portpb"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// Construct the application logger.
	log, err := foundation.NewLogger("CLIENT")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()

	code := run(log)
	os.Exit(code)
}

func run(log *zap.SugaredLogger) int {
	var wait time.Duration = 5 * time.Second

	// Load environment variables
	cfg := &api.Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Error(err)
		return 1
	}

	// GRPC connection
	ctxGRPC, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	address := cfg.GRPCHost + ":" + cfg.GRPCPort
	log.Infof("Starting client GRPC connection ", address)

	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	grpcClientConn, err := grpc.DialContext(ctxGRPC, address, opts...)
	if err != nil {
		log.Error("did not connect: %s %v", address, err)
		return 1
	}

	// add GRPC to the PortServiceClient
	cfg.PSC = portpb.NewPortServiceClient(grpcClientConn)

	// initialize the Core
	core := api.CoreConfig{
		Port: port.NewCore(log, cfg.PSC),
	}

	// Create the API handlers, routes and middlewares
	a := api.New()
	a.Routes(core, cfg)
	a.Mid(cfg)

	srv := &http.Server{
		Addr:         "0.0.0.0:" + cfg.APIPort,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      a.Handler(),
	}

	log.Infof("Running client api at port:", cfg.APIPort)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error(err)
		}
	}()

	// the graceful shutdown
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	log.Error("Closing GRPC Connection")
	defer grpcClientConn.Close()

	if err := srv.Shutdown(ctx); err != nil {
		log.Errorf("Error trying to shutting down: %s", err)
	}
	log.Error("shutting down")

	return 0

}
