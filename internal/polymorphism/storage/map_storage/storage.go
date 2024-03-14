package map_storage

import "test/internal/polymorphism/storage"

type Storage struct {
	store map[string]string
}

func NewStorage() *Storage {
	return &Storage{
		store: make(map[string]string, 100),
	}
}

func (s *Storage) GetValue(key string) (string, error) {
	value, found := s.store[key]
	if !found {
		return value, storage.ErrNotFound
	}

	return value, nil
}

func (s *Storage) SavePair(key, value string) error {
	s.store[key] = value
	return nil
}

func (s *Storage) Close() error {
	return nil
}
