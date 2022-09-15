package persistence

import (
	"fmt"

	"github.com/eduardobcolombo/learning-grpc/cmd/server/domain/entity"
	"github.com/eduardobcolombo/learning-grpc/cmd/server/domain/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Host     string `envconfig:"DB_HOST" default:"postgres"`
	Password string `envconfig:"DB_PASSWORD" default:"passwd"`
	User     string `envconfig:"DB_USER" default:"user"`
	Name     string `envconfig:"DB_NAME" default:"db"`
	Port     string `envconfig:"DB_PORT" default:"5432"`
}

type Repositories struct {
	Port repository.PortRepository
	db   *gorm.DB
}

func NewRepositories(DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repositories{
		Port: NewPortRepository(db),
		db:   db,
	}, nil
}

func (s *Repositories) Close() error {
	sqlDB, _ := s.db.DB()

	return sqlDB.Close()
}

func (s *Repositories) Automigrate() error {
	return s.db.AutoMigrate(&entity.Port{}, &entity.Alias{}, &entity.Coordinate{}, &entity.Region{}, &entity.Unloc{})
}
