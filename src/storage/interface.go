package storage

// Storage ...
type Storage interface {
	Get(key string) (string, error)
	Put(key, value string) error
	Delete(key string) error
}
