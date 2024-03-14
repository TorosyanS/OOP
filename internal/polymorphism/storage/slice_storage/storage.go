package slice_storage

import "test/internal/polymorphism/storage"

type pair struct {
	key   string
	value string
}

type Storage struct {
	store []pair
}

func NewStorage() *Storage {
	return &Storage{
		store: make([]pair, 0, 100),
	}
}

func (s *Storage) GetValue(key string) (string, error) {
	for _, p := range s.store {
		if p.key == key {
			return p.value, nil
		}
	}

	return "", storage.ErrNotFound
}

func (s *Storage) SavePair(key, value string) error {
	s.store = append(s.store, pair{
		key:   key,
		value: value,
	})
	return nil
}

func (s *Storage) Close() error {
	return nil
}
