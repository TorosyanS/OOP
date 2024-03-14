package storage

type Storage interface {
	GetValue(key string) (string, error)
	SavePair(key, value string) error
	Close() error
}
