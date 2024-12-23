package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/eduardobcolombo/learning-grpc/cmd/server/app"
	"github.com/eduardobcolombo/learning-grpc/cmd/server/core/port"
	"github.com/eduardobcolombo/learning-grpc/cmd/server/core/port/store"
	"github.com/eduardobcolombo/learning-grpc/cmd/server/handlers"
	"github.com/eduardobcolombo/learning-grpc/foundation"
	"github.com/eduardobcolombo/learning-grpc/foundation/sqlDB"
	sqlxdb "github.com/eduardobcolombo/learning-grpc/foundation/sqlxDB"
	"github.com/eduardobcolombo/learning-grpc/portpb"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Config struct {
	GRPC        GRPC
	DBConfig    *sqlDB.DBConfig
	DBxConfig   *sqlxdb.DBConfig
	Environment string `envconfig:"ENVIRONMENT" default:"production"`
	Log         foundation.LoggerConfig
	TLS         bool `envconfig:"TLS" default:"false"`
}

type GRPC struct {
	Host string `envconfig:"GRPC_HOST" default:"0.0.0.0"`
	Port string `envconfig:"GRPC_PORT" default:"50053"`
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Construct the application logger.
	logger, err := foundation.NewLogger(&cfg.Log)
	if err != nil {
		fmt.Printf("initialize logger: %v\n", err)
		return err
	}
	logger.Z = logger.Z.With(zap.String("env", cfg.Environment))

	logger.Infof("Server is running...\n")

	address := fmt.Sprintf("%s:%s", cfg.GRPC.Host, cfg.GRPC.Port)
	logger.Infof("-----> %s\n", address)
	list, err := net.Listen("tcp", address)
	if err != nil {
		logger.Errorf("Failed to listen: %v", err)
		return err
	}

	// I'm leaving the DBConfig commented for demo purpose.
	// This is the implementation for the Postgres DB uses GORM.
	// pg, err := sqlDB.New(cfg.DBConfig)
	// if err != nil {
	// 	logger.Errorf("error initializating the DB: %v", err)
	// 	return err
	// }
	// defer pg.Close()

	// if err := pg.AutoMigrate(context.Background(), &entity.Port{}); err != nil {
	// 	logger.Errorf("error running migrations: %v", err)
	// 	return err
	// }

	pgx, err := sqlxdb.New(cfg.DBxConfig)
	if err != nil {
		logger.Errorf("error initializating the DB: ", err)
		return err
	}
	defer pgx.Close()

	if err := sqlxdb.Migrate(cfg.DBxConfig); err != nil {
		logger.Errorf("error running migrations: ", err)
		return err
	}

	srv := app.Server{
		Log: logger,
		// Select the store to be used.
		// Port: *handlers.NewPort(port.NewCore(db.NewStore(pg.DB))),
		// Port: *handlers.NewPort(port.NewCore(memory.NewStore())),
		Port: *handlers.NewPort(port.NewCore(store.NewStore(pgx))),
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	portpb.RegisterPortServiceServer(s, &srv)
	reflection.Register(s)

	// Prepare to shutdown.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start the server.
	go func() {
		if err := s.Serve(list); err != nil {
			logger.Error("failed to serve: ", err)
		}
	}()

	// Wait for the shutdown signal.
	<-shutdown

	logger.Info("Stopping the server\n")
	s.GracefulStop()
	logger.Info("Closing the listener")
	list.Close()
	logger.Info("Shutdown")

	return nil
}
