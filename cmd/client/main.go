package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/eduardobcolombo/learning-grpc/cmd/client/api"
	"github.com/eduardobcolombo/learning-grpc/cmd/client/handlers"
	"github.com/eduardobcolombo/learning-grpc/foundation"
	"github.com/eduardobcolombo/learning-grpc/portpb"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var wait time.Duration = 5 * time.Second

	// load environment variables.
	cfg := &api.Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// start the application logger.
	logger, err := foundation.NewLogger(&cfg.Log)
	if err != nil {
		fmt.Printf("initialize logger: %v\n", err)
		return err
	}
	logger.Z = logger.Z.With(zap.String("env", cfg.Environment))

	address := fmt.Sprintf("%s:%s", cfg.GRPCHost, cfg.GRPCPort)
	logger.Infof("Starting client GRPC connection ", address)

	// start the GRPC configuration.
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// start the GRPC connection.
	grpcClientConn, err := grpc.NewClient(address, opts...)
	if err != nil {
		logger.Errorf("did not connect: %s %v", address, err)
		return err
	}

	defer grpcClientConn.Close()

	// add GRPC conn to the PortServiceClient.
	cfg.PSC = portpb.NewPortServiceClient(grpcClientConn)

	// initialize the Core.
	// core := api.CoreConfig{
	// 	Port: port.NewCore(logger, cfg.PSC),
	// }
	core := api.NewCoreConfig(logger, cfg.PSC)

	// the API handlers, routes and middlewares.
	a := api.New()
	handlerPort := handlers.NewHandler(core.Port, logger)

	a.Routes(handlerPort)
	a.Mid(cfg)

	srv := &http.Server{
		Addr:         "0.0.0.0:" + cfg.APIPort,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      a.Router,
	}

	logger.Infof("Running client api at port:", cfg.APIPort)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Error(err)
		}
	}()

	// the graceful shutdown.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Errorf("Error trying to shutting down: %s", err)
		// If the graceful shutdown do not work, it will close anyway.
		srv.Close()
	}
	logger.Error("shutting down")

	return nil
}
