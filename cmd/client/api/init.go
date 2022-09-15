package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/eduardobcolombo/learning-grpc/internal/pkg/port"
	"github.com/eduardobcolombo/learning-grpc/internal/pkg/portpb"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func Initialize(log *zap.SugaredLogger) int {
	var wait time.Duration
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Error(err)
		return 1
	}

	grpcClientConn, err := grpcInit(cfg)
	if err != nil {
		log.Error(err)
		return 1
	}

	cfg.psc = portpb.NewPortServiceClient(grpcClientConn)

	core := CoreConfig{
		Port: port.NewCore(log, cfg.psc),
	}

	api := newAPI()
	api.routes(core, cfg)
	api.mid(cfg)

	srv := &http.Server{
		Addr:         "0.0.0.0:" + cfg.APIPort,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      api.Handler(),
	}

	fmt.Println("Running client api at port:", cfg.APIPort)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error(err)
		}
	}()

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

func grpcInit(cfg *Config) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	address := cfg.GRPCHost + ":" + cfg.GRPCPort
	fmt.Println("Starting client GRPC connection ", address)

	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// I'm using cfg.tls always false, but here is the implementation if we would like to make it tls based
	if cfg.tls {
		cFile := "ADD_THE_CERTIFICATE_PATH_HERE"
		crds, err := credentials.NewClientTLSFromFile(cFile, "")
		if err != nil {
			log.Printf("Error loading certificate: %v", err)
			return nil, err
		}
		opts = append(opts, grpc.WithTransportCredentials(crds))
	}

	cc, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		log.Printf("did not connect: %s %v", address, err)
		return nil, err
	}

	return cc, nil
}
