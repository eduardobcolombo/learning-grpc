package persistence

import (
	"fmt"

	"github.com/eduardobcolombo/learning-grpc/server/domain/entity"
	"github.com/eduardobcolombo/learning-grpc/server/domain/repository"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Repositories struct {
	Port repository.PortRepository
	db   *gorm.DB
}

func NewRepositories(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(Dbdriver, DBURL)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &Repositories{
		Port: NewPortRepository(db),
		db:   db,
	}, nil
}

func (s *Repositories) Close() error {
	return s.db.Close()
}

func (s *Repositories) Automigrate() error {
	return s.db.AutoMigrate(&entity.Port{}).Error
}
