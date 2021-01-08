package storage

import (
	"errors"
	"sync"
)

// InMemory ...
type InMemory struct {
	mu    sync.RWMutex
	store map[string]string
}

// InitInMemory ...
func InitInMemory() Storage {
	return &InMemory{
		store: map[string]string{},
	}
}

// Get ...
func (st *InMemory) Get(key string) (string, error) {
	st.mu.RLock()
	defer st.mu.RUnlock()
	value, ok := st.store[key]
	if !ok {
		return "", errors.New("no data")
	}
	return value, nil
}

// Put ...
func (st *InMemory) Put(key, value string) error {
	st.mu.Lock()
	defer st.mu.Unlock()
	st.store[key] = value
	return nil
}

// Delete ...
func (st *InMemory) Delete(key string) error {
	st.mu.Lock()
	defer st.mu.Unlock()
	_, ok := st.store[key]
	if ok {
		delete(st.store, key)
	}
	return nil
}
