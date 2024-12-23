package db

import (
	"context"
	"errors"

	"github.com/eduardobcolombo/learning-grpc/cmd/server/domain/entity"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) Create(ctx context.Context, port entity.Port) error {
	return s.db.WithContext(ctx).Create(port).Error
}

func (s *Store) Update(ctx context.Context, port entity.Port) error {
	return s.db.WithContext(ctx).Save(&port).Error
}

func (s *Store) Delete(ctx context.Context, id uint) error {
	return s.db.WithContext(ctx).Delete(entity.Port{}, id).Error
}

func (s *Store) GetByID(ctx context.Context, id uint) (*entity.Port, error) {
	var port entity.Port
	if err := s.db.WithContext(ctx).Where(entity.Port{ID: id}).Find(&port).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, entity.ErrPortNotFound
		}
		return nil, err
	}

	return &port, nil
}

func (s *Store) GetByUnloc(ctx context.Context, unloc string) (*entity.Port, error) {
	var port entity.Port
	if err := s.db.WithContext(ctx).Where(entity.Port{Unlocs: unloc}).Find(&port).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, entity.ErrPortNotFound
		}
		return nil, err
	}

	return &port, nil
}

func (s *Store) GetAll(ctx context.Context) ([]entity.Port, error) {
	var ports []entity.Port

	if err := s.db.WithContext(ctx).Find(&ports).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, entity.ErrPortNotFound
		}
		return nil, err
	}

	return ports, nil
}
