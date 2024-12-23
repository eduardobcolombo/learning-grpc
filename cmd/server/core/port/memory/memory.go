package memory

import (
	"context"
	"sync"

	"github.com/eduardobcolombo/learning-grpc/cmd/server/domain/entity"
)

type Store struct {
	store []entity.Port
	sync.Mutex
}

func NewStore() *Store {
	return &Store{
		store: []entity.Port{},
	}
}

func (s *Store) Create(ctx context.Context, port entity.Port) error {
	s.Lock()
	defer s.Unlock()

	port.ID = s.getNextID()
	s.store = append(s.store, port)

	return nil
}

func (s *Store) Update(ctx context.Context, port entity.Port) error {
	s.Lock()
	defer s.Unlock()

	id := port.ID

	for i := 0; i < len(s.store); i++ {
		if s.store[i].ID == id {
			s.store[i] = port
			return nil
		}
	}

	return entity.ErrPortNotFound
}

func (s *Store) Delete(ctx context.Context, id uint) error {
	s.Lock()
	defer s.Unlock()

	for i := 0; i < len(s.store); i++ {
		if s.store[i].ID == id {
			s.store = append(s.store[:i], s.store[i+1:]...)
			return nil
		}
	}

	return entity.ErrPortNotFound
}

func (s *Store) GetByID(ctx context.Context, id uint) (*entity.Port, error) {
	s.Lock()
	defer s.Unlock()

	for _, v := range s.store {
		if v.ID == id {
			return &v, nil
		}
	}

	return &entity.Port{}, entity.ErrPortNotFound
}

func (s *Store) GetByUnloc(ctx context.Context, unloc string) (*entity.Port, error) {
	s.Lock()
	defer s.Unlock()

	for _, v := range s.store {
		if v.Unlocs == unloc {
			return &v, nil
		}
	}

	return &entity.Port{}, entity.ErrPortNotFound
}

func (s *Store) GetAll(ctx context.Context) ([]entity.Port, error) {
	return s.store, nil
}

func (s *Store) getNextID() uint {
	var id uint

	if len(s.store) == 0 {
		return 1
	}

	for _, value := range s.store {
		if value.ID > id {
			id = value.ID
		}
	}

	return id + 1
}
