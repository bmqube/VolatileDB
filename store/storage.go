package store

import "sync"

type Storage struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewStorage() *Storage {
	return &Storage{data: make(map[string]string)}
}

func (store *Storage) Get(key string) (string, bool) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	val, ok := store.data[key]

	return val, ok
}

func (store *Storage) Set(key string, val string) {
	store.mu.Lock()
	defer store.mu.Unlock()
	store.data[key] = val
}
