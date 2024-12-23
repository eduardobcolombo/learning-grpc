package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/eduardobcolombo/learning-grpc/cmd/server/domain/entity"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{db: db}
}

func (s Store) Create(ctx context.Context, port entity.Port) error {
	res, err := s.db.ExecContext(ctx,
		`INSERT INTO ports (name, city, country, alias, regions, latitude, longitude, province, timezone, unlocs, code, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
		port.Name, port.City, port.Country, port.Alias, port.Regions, port.Latitude, port.Longitude, port.Province, port.Timezone, port.Unlocs, port.Code, time.Now())
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return entity.ErrOnInsert
	}

	return err
}

func (s Store) Update(ctx context.Context, port entity.Port) error {

	now := time.Now()
	res, err := s.db.ExecContext(ctx,
		`UPDATE ports SET 
			name = $1, city = $2, country = $3, alias = $4, regions = $5, latitude = $6, longitude = $7, province = $8, timezone = $9, unlocs = $10, code = $11, updated_at = $13
		WHERE id = $12;`,
		port.Name, port.City, port.Country, port.Alias, port.Regions, port.Latitude, port.Longitude, port.Province, port.Timezone, port.Unlocs, port.Code, port.ID, now)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return entity.ErrOnUpdate
	}

	return err
}

func (s Store) Delete(ctx context.Context, id uint) error {
	res, err := s.db.ExecContext(ctx,
		`DELETE FROM ports WHERE id = $1;`,
		id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return entity.ErrPortNotFound
	}

	return err
}

func (s Store) GetByID(ctx context.Context, id uint) (*entity.Port, error) {
	var port entity.Port

	if err := s.db.GetContext(ctx, &port,
		`SELECT * FROM ports WHERE id = $1;`,
		id,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &entity.Port{}, entity.ErrPortNotFound
		}
		return &entity.Port{}, err
	}

	return &port, nil
}

func (s Store) GetByUnloc(ctx context.Context, unloc string) (*entity.Port, error) {
	var port entity.Port

	if err := s.db.GetContext(ctx, &port,
		`SELECT * FROM ports WHERE unlocs = $1;`,
		unloc,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &entity.Port{}, entity.ErrPortNotFound
		}
		return &entity.Port{}, err
	}

	return &port, nil
}

func (s Store) GetAll(ctx context.Context) ([]entity.Port, error) {
	ports := []entity.Port{}

	if err := s.db.SelectContext(ctx, &ports,
		`SELECT * FROM ports `,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.ErrPortNotFound
		}
		return nil, err
	}

	return ports, nil
}
