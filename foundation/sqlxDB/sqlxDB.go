package sqlxdb

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/kulado/sqlxmigrate"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host           string `envconfig:"DB_HOST" default:"postgres"`
	Password       string `envconfig:"DB_PASSWORD" default:"passwd"`
	User           string `envconfig:"DB_USER" default:"user"`
	Name           string `envconfig:"DB_NAME" default:"db"`
	Port           string `envconfig:"DB_PORT" default:"5432"`
	MigrationsPath string `envconfig:"MIGRATIONS_PATH" default:"./migrations"`
}

type Migration struct {
	ID   string
	UP   string
	DOWN string
}

// New open a db connection
func New(cfg *DBConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", cfg.Host, cfg.Port, cfg.User, cfg.Name, cfg.Password)

	pg, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return pg, nil
}

func Migrate(cfg *DBConfig) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", cfg.Host, cfg.Port, cfg.User, cfg.Name, cfg.Password)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return fmt.Errorf("migrate : Register DB : %v", err)
	}
	defer db.Close()

	migrations, err := GetMigrations(cfg.MigrationsPath)
	if err != nil {
		return fmt.Errorf("could not get migrations: %v", err)
	}

	sqlxMigrations := make([]*sqlxmigrate.Migration, 0)
	for _, m := range migrations {
		sqlxMigrations = append(sqlxMigrations, &sqlxmigrate.Migration{
			ID: m.ID,
			Migrate: func(tx *sql.Tx) error {
				_, err = tx.Exec(m.UP)
				return err
			},
			Rollback: func(tx *sql.Tx) error {
				_, err = tx.Exec(m.DOWN)
				return err
			},
		})
	}

	if len(sqlxMigrations) > 0 {
		sqlxPrep := sqlxmigrate.New(db, sqlxmigrate.DefaultOptions, sqlxMigrations)

		if err = sqlxPrep.Migrate(); err != nil {
			return fmt.Errorf("could not migrate: %v", err)

		}
	}

	fmt.Println("Migration did run successfully.")

	return nil
}

func GetMigrations(path string) ([]Migration, error) {
	// read the migrations from the files in the path
	// the prefix until the first underline will be the ID
	// the files ended with .up.sql will be the UP migration
	// the files ended with .down.sql will be the DOWN migration

	upPrefix := ".up.sql"
	downPrefix := ".down.sql"

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	migrations := make([]Migration, 0)

	var up, down []byte

	for i := 0; i < len(files); i++ {
		up = nil

		file := files[i]

		if file.Name()[len(file.Name())-len(downPrefix):] != downPrefix {
			continue
		}

		down, err = os.ReadFile(path + "/" + file.Name())
		if err != nil {
			return nil, err
		}

		if i+1 > len(files) {
			return nil, fmt.Errorf("missing up migration for %s", file.Name())
		}

		nextFile := files[i+1]

		if file.Name()[:len(file.Name())-len(downPrefix)] == nextFile.Name()[:len(nextFile.Name())-len(upPrefix)] {
			i = i + 1
			up, err = os.ReadFile(path + "/" + nextFile.Name())
			if err != nil {
				return nil, err
			}
		}

		migrations = append(migrations, Migration{
			ID:   file.Name()[:6],
			UP:   string(up),
			DOWN: string(down),
		})
	}

	return migrations, nil
}
