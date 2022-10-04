package sqlDB

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

type DBConfig struct {
	Host     string `envconfig:"DB_HOST" default:"postgres"`
	Password string `envconfig:"DB_PASSWORD" default:"passwd"`
	User     string `envconfig:"DB_USER" default:"user"`
	Name     string `envconfig:"DB_NAME" default:"db"`
	Port     string `envconfig:"DB_PORT" default:"5432"`
}

// New open a db connection
func New(cfg *DBConfig) (*DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", cfg.Host, cfg.Port, cfg.User, cfg.Name, cfg.Password)

	pg, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DB{DB: pg}, nil
}

// Automigrate run migrations for the db
func (d *DB) Automigrate(ent ...interface{}) error {
	return d.DB.AutoMigrate(ent...)
}

// Close close the db connection
func (d *DB) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
