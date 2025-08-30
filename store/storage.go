package store

import (
	"errors"
	"strconv"
	"sync"
)

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

func (store *Storage) Incr(key string) (int64, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	val, ok := store.data[key]
	if !ok {
		store.data[key] = "0"
		val = "0"
	}

	intVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, errors.New("ERR value is not an integer or out of range")
	}

	intVal++
	store.data[key] = strconv.FormatInt(intVal, 10)

	return intVal, nil
}

func (store *Storage) Decr(key string) (int64, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	val, ok := store.data[key]
	if !ok {
		store.data[key] = "0"
		val = "0"
	}

	intVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, errors.New("ERR value is not an integer or out of range")
	}

	intVal--
	store.data[key] = strconv.FormatInt(intVal, 10)

	return intVal, nil
}

func (store *Storage) Del(key string) bool {
	store.mu.Lock()
	defer store.mu.Unlock()
	_, ok := store.data[key]
	if ok {
		delete(store.data, key)
	}
	return ok
}
