package file_storage

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"test/internal/polymorphism/storage"
)

const storageFileName = "storage.txt"

type Storage struct {
	file         *os.File
	inMemStorage map[string]string
}

func NewStorage() (*Storage, error) {
	file, err := os.OpenFile(storageFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0o644)
	if err != nil {
		return nil, err
	}
	store := &Storage{
		inMemStorage: make(map[string]string, 100),
		file:         file,
	}
	err = store.loadData()
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (s *Storage) GetValue(key string) (string, error) {
	value, found := s.inMemStorage[key]
	if !found {
		return value, storage.ErrNotFound
	}

	return value, nil
}

func (s *Storage) SavePair(key, value string) error {
	_, err := s.file.WriteString(fmt.Sprintf("%s|%s\n", key, value))
	if err != nil {
		return err
	}
	s.inMemStorage[key] = value

	return nil
}

func (s *Storage) loadData() error {
	scanner := bufio.NewScanner(s.file)
	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), "|")
		if len(pair) != 2 {
			continue
		}
		s.inMemStorage[pair[0]] = pair[1]
	}
	return scanner.Err()
}

func (s *Storage) Close() error {
	return s.file.Close()
}
