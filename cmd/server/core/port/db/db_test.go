package db

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/eduardobcolombo/learning-grpc/cmd/server/domain/entity"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// These tests uses testify pkg to show up how to use it

func TestStore_GetByUnloc(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlxDB.DB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	require.NoError(t, err)

	stg := NewStore(gormDB)

	tests := []struct {
		name        string
		unloc       string
		mock        func()
		wantErr     bool
		expectedErr error
	}{
		{
			name:  "Connection error",
			unloc: "unlocs",
			mock: func() {
				mock.ExpectQuery("SELECT \\* FROM \"ports\" WHERE \"ports\".\"unlocs\" = \\$1").WithArgs("unlocs").WillReturnError(errors.New("connection error"))
			},
			wantErr:     true,
			expectedErr: errors.New("connection error"),
		},
		{
			name:  "No rows found",
			unloc: "unlocs",
			mock: func() {
				mock.ExpectQuery("SELECT \\* FROM \"ports\" WHERE \"ports\".\"unlocs\" = \\$1").WithArgs("unlocs").WillReturnError(entity.ErrPortNotFound)
			},
			wantErr:     true,
			expectedErr: entity.ErrPortNotFound,
		},
		{
			name:  "No rows found",
			unloc: "unlocs",
			mock: func() {
				mock.ExpectQuery("SELECT \\* FROM \"ports\" WHERE \"ports\".\"unlocs\" = \\$1").WithArgs("unlocs").WillReturnError(gorm.ErrRecordNotFound)
			},
			wantErr:     true,
			expectedErr: entity.ErrPortNotFound,
		},
		{
			name:  "Successful retrieval",
			unloc: "unlocs",
			mock: func() {
				mock.ExpectQuery("SELECT \\* FROM \"ports\" WHERE \"ports\".\"unlocs\" = \\$1").
					WithArgs("unlocs").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "city", "country", "alias", "regions", "latitude", "longitude", "province", "timezone", "unlocs", "code"}).
						AddRow(1, "name", "city", "country", "alias", "regions", 12.123, 12.123, "province", "timezone", "unlocs", "code"))
			},
			wantErr:     false,
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			port, err := stg.GetByUnloc(ctx, tt.unloc)
			if tt.wantErr {
				require.Error(t, err)
				require.Equal(t, tt.expectedErr, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, port)
				require.Equal(t, tt.unloc, port.Unlocs)
			}
		})
	}

	// make sure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestStore_GetByID(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlxDB.DB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	require.NoError(t, err)

	stg := NewStore(gormDB)

	tests := []struct {
		name        string
		id          uint
		mock        func()
		wantErr     bool
		expectedErr error
	}{
		{
			name: "Connection error",
			id:   999,
			mock: func() {
				mock.ExpectQuery("SELECT \\* FROM \"ports\" WHERE \"ports\".\"id\" = \\$1").WithArgs(999).WillReturnError(errors.New("connection error"))
			},
			wantErr:     true,
			expectedErr: errors.New("connection error"),
		},
		{
			name: "No rows found",
			id:   999,
			mock: func() {
				mock.ExpectQuery("SELECT \\* FROM \"ports\" WHERE \"ports\".\"id\" = \\$1").WithArgs(999).WillReturnError(entity.ErrPortNotFound)
			},
			wantErr:     true,
			expectedErr: entity.ErrPortNotFound,
		},
		{
			name: "No rows found",
			id:   999,
			mock: func() {
				mock.ExpectQuery("SELECT \\* FROM \"ports\" WHERE \"ports\".\"id\" = \\$1").WithArgs(999).WillReturnError(gorm.ErrRecordNotFound)
			},
			wantErr:     true,
			expectedErr: entity.ErrPortNotFound,
		},
		{
			name: "Successful retrieval",
			id:   999,
			mock: func() {
				mock.ExpectQuery("SELECT \\* FROM \"ports\" WHERE \"ports\".\"id\" = \\$1").
					WithArgs(999).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "city", "country", "alias", "regions", "latitude", "longitude", "province", "timezone", "unlocs", "code"}).
						AddRow(999, "name", "city", "country", "alias", "regions", 12.123, 12.123, "province", "timezone", "unloc", "code"))
			},
			wantErr:     false,
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			port, err := stg.GetByID(ctx, tt.id)
			if tt.wantErr {
				require.Error(t, err)
				require.Equal(t, tt.expectedErr, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, port)
				require.Equal(t, tt.id, port.ID)
			}
		})
	}

	// make sure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestStore_GetAll(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlxDB.DB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	require.NoError(t, err)

	stg := NewStore(gormDB)

	tests := []struct {
		name        string
		mock        func()
		wantErr     bool
		expectedErr error
		expectedLen int
	}{
		{
			name: "No rows found",
			mock: func() {
				mock.ExpectQuery("SELECT \\* FROM \"ports\" ").WillReturnError(entity.ErrPortNotFound)
			},
			wantErr:     true,
			expectedErr: entity.ErrPortNotFound,
			expectedLen: 0,
		},
		{
			name: "No rows found",
			mock: func() {
				mock.ExpectQuery("SELECT \\* FROM \"ports\" ").WillReturnError(gorm.ErrRecordNotFound)
			},
			wantErr:     true,
			expectedErr: entity.ErrPortNotFound,
			expectedLen: 0,
		},
		{
			name: "Successful GetAll",
			mock: func() {
				mock.ExpectQuery("SELECT \\* FROM \"ports\" ").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "city", "country", "alias", "regions", "latitude", "longitude", "province", "timezone", "unlocs", "code"}).
						AddRow(1, "name", "city", "country", "alias", "regions", 12.123, 12.123, "province", "timezone", "unlocs", "code").
						AddRow(2, "name 2", "city 2", "country 2", "alias 2", "regions 2", 22.123, 22.123, "province 2", "timezone 2", "unlocs", "code 2"))
			},
			wantErr:     false,
			expectedErr: nil,
			expectedLen: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			ports, err := stg.GetAll(ctx)
			if tt.wantErr {
				require.Error(t, err)
				require.Equal(t, tt.expectedErr, err)
				require.Equal(t, []entity.Port(nil), ports)
			} else {
				require.NoError(t, err)
				require.NotNil(t, ports)
				require.Equal(t, tt.expectedLen, len(ports))
			}
		})
	}

	// make sure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
