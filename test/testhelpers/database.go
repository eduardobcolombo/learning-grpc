package testhelpers

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type DBCfg struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func buildpgContainer(ctx context.Context, dbName, dbUser, dbPassword string) (testcontainers.Container, error) {
	return postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:16-alpine"),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
}

func DatabaseContainer() (context.Context, testcontainers.Container, *DBCfg) {
	ctx := context.Background()

	dbName := uuid.New().String()
	dbUser := uuid.New().String()
	dbPassword := uuid.New().String()

	pgContainer, err := buildpgContainer(ctx, dbName, dbUser, dbPassword)
	if err != nil {
		log.Fatalf("failed to build postgres container: %s", err)
	}

	port, err := pgContainer.MappedPort(ctx, "5432/tcp")
	if err != nil {
		log.Fatalf("failed to get mapped port: %s", err)
	}
	host, err := pgContainer.Host(ctx)
	if err != nil {
		log.Fatalf("failed to get host: %s", err)
	}

	dbCfg := &DBCfg{
		Host:     host,
		Port:     port.Port(),
		User:     dbUser,
		Password: dbPassword,
		Name:     dbName,
	}

	return ctx, pgContainer, dbCfg
}
