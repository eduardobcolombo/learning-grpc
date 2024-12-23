package sqlxdb

import (
	"log"
	"testing"

	"github.com/eduardobcolombo/learning-grpc/test/testhelpers"
)

func TestSQLDBx(t *testing.T) {

	ctxDB, pgContainer, dbCfg := testhelpers.DatabaseContainer()

	pgCfg := &DBConfig{
		Host:           dbCfg.Host,
		Port:           dbCfg.Port,
		User:           dbCfg.User,
		Password:       dbCfg.Password,
		Name:           dbCfg.Name,
		MigrationsPath: "../../migrations",
	}

	t.Run("should connect to the database", func(t *testing.T) {

		db, err := New(pgCfg)
		if err != nil {
			t.Errorf("should not return error openning the DB connection, got: %v", err)
		}

		if err := Migrate(pgCfg); err != nil {
			t.Errorf("should not return error running migrations, got: %v", err)
		}

		if err := db.Close(); err != nil {
			t.Errorf("should not return error closing the DB connection, got: %v", err)
		}

		t.Cleanup(func() {
			if err := pgContainer.Terminate(ctxDB); err != nil {
				log.Fatalf("failed to terminate container: %s", err)
			}
		})
	})
}
