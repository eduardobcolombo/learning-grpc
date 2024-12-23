package store

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/eduardobcolombo/learning-grpc/cmd/server/domain/entity"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestStore(t *testing.T) {

	ctx := context.Background()
	db, mock, errDB := sqlmock.New()
	require.NoError(t, errDB)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	stg := NewStore(sqlxDB)

	// GetByID
	var id uint = 1
	mock.ExpectQuery("SELECT \\* FROM ports WHERE ").WithArgs(id).WillReturnError(sql.ErrConnDone)
	mock.ExpectQuery("SELECT \\* FROM ports WHERE ").WithArgs(id).WillReturnError(sql.ErrNoRows)
	mock.ExpectQuery("SELECT \\* FROM ports WHERE ").WithArgs(id).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "city", "country", "alias", "regions", "latitude", "longitude", "province", "timezone", "unlocs", "code"}).AddRow(id, "name", "city", "country", "alias", "regions", 12.123, 12.123, "province", "timezone", "unlocs", "code"))

	_, err := stg.GetByID(ctx, id)
	require.ErrorAs(t, sql.ErrConnDone, &err)

	_, err = stg.GetByID(ctx, id)
	require.ErrorAs(t, sql.ErrNoRows, &err)

	data, err := stg.GetByID(ctx, id)
	require.Nil(t, err)
	require.Equal(t, id, data.ID)

	// GetByUnloc
	var unloc string = "unlocs"
	mock.ExpectQuery("SELECT \\* FROM ports WHERE ").WithArgs(unloc).WillReturnError(sql.ErrConnDone)
	mock.ExpectQuery("SELECT \\* FROM ports WHERE ").WithArgs(unloc).WillReturnError(sql.ErrNoRows)
	mock.ExpectQuery("SELECT \\* FROM ports WHERE ").WithArgs(unloc).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "city", "country", "alias", "regions", "latitude", "longitude", "province", "timezone", "unlocs", "code"}).AddRow(1, "name", "city", "country", "alias", "regions", 12.123, 12.123, "province", "timezone", "unlocs", "code"))

	_, err = stg.GetByUnloc(ctx, unloc)
	require.ErrorAs(t, sql.ErrConnDone, &err)

	_, err = stg.GetByUnloc(ctx, unloc)
	require.ErrorAs(t, sql.ErrNoRows, &err)

	data, err = stg.GetByUnloc(ctx, unloc)
	require.Nil(t, err)
	require.Equal(t, unloc, data.Unlocs)

	// GetAll
	mock.ExpectQuery("SELECT \\* FROM ports ").WillReturnError(sql.ErrConnDone)
	mock.ExpectQuery("SELECT \\* FROM ports ").WillReturnError(sql.ErrNoRows)
	mock.ExpectQuery("SELECT \\* FROM ports ").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "city", "country", "alias", "regions", "latitude", "longitude", "province", "timezone", "unlocs", "code"}).
		AddRow(1, "name", "city", "country", "alias", "regions", 12.123, 12.123, "province", "timezone", "unlocs", "code").
		AddRow(2, "name", "city", "country", "alias", "regions", 12.123, 12.123, "province", "timezone", "unlocs", "code"))

	_, err = stg.GetAll(ctx)
	require.ErrorAs(t, sql.ErrConnDone, &err)

	_, err = stg.GetAll(ctx)
	require.ErrorAs(t, sql.ErrNoRows, &err)

	ports, err := stg.GetAll(ctx)
	require.Nil(t, err)
	require.Equal(t, 2, len(ports))

	// Create
	mock.ExpectExec("INSERT INTO port").WillReturnError(sql.ErrConnDone)
	mock.ExpectExec("INSERT INTO port").WillReturnError(sql.ErrNoRows)
	mock.ExpectExec("INSERT INTO port").WillReturnResult(sqlmock.NewResult(1, 1))

	err = stg.Create(ctx, entity.Port{})
	require.ErrorAs(t, sql.ErrConnDone, &err)

	err = stg.Create(ctx, entity.Port{})
	require.ErrorAs(t, sql.ErrNoRows, &err)

	err = stg.Create(ctx, entity.Port{})
	require.Nil(t, err)

	// Update
	mock.ExpectExec("UPDATE port").WillReturnError(sql.ErrConnDone)
	mock.ExpectExec("UPDATE port").WillReturnError(sql.ErrNoRows)
	mock.ExpectExec("UPDATE port").WillReturnResult(sqlmock.NewResult(1, 1))

	err = stg.Update(ctx, entity.Port{})
	require.ErrorAs(t, sql.ErrConnDone, &err)

	err = stg.Update(ctx, entity.Port{})
	require.ErrorAs(t, sql.ErrNoRows, &err)

	err = stg.Update(ctx, entity.Port{ID: 1, Name: "Updated"})
	require.Nil(t, err)

	// Delete
	mock.ExpectExec("DELETE FROM ports").WillReturnError(sql.ErrConnDone)
	mock.ExpectExec("DELETE FROM ports").WillReturnError(sql.ErrNoRows)
	mock.ExpectExec("DELETE FROM ports").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	err = stg.Delete(ctx, 1)
	require.ErrorAs(t, sql.ErrConnDone, &err)

	err = stg.Delete(ctx, 1)
	require.ErrorAs(t, sql.ErrNoRows, &err)

	err = stg.Delete(ctx, 1)
	require.Nil(t, err)

	// make sure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
