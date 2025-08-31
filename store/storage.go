package store

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

type Storage struct {
	mu   sync.RWMutex
	data map[string]Entry
}

type Entry struct {
	Value     string
	ExpiresAt time.Time
}

func NewStorage() *Storage {
	s := &Storage{
		data: make(map[string]Entry),
	}
	s.startCleaner(1 * time.Second)
	return s
}

func (s *Storage) startCleaner(interval time.Duration) {
	go func() {
		for {
			time.Sleep(interval)
			now := time.Now()

			var deletableKeys []string

			s.mu.RLock()
			for k, e := range s.data {
				if !e.ExpiresAt.IsZero() && now.After(e.ExpiresAt) {
					deletableKeys = append(deletableKeys, k)
				}
			}
			s.mu.RUnlock()

			if len(deletableKeys) > 0 {
				s.mu.Lock()

				for _, k := range deletableKeys {
					delete(s.data, k)
				}

				s.mu.Unlock()
			}
		}
	}()
}

func (store *Storage) Get(key string) (string, bool) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	entry, ok := store.data[key]

	if ok && !entry.ExpiresAt.IsZero() && time.Now().After(entry.ExpiresAt) {
		store.mu.Lock()
		delete(store.data, key)
		store.mu.Unlock()
		ok = false
	}

	if !ok {
		return "", false
	}

	return entry.Value, true
}

func (store *Storage) Set(key string, val string, ttl time.Duration) {
	store.mu.Lock()
	defer store.mu.Unlock()

	var expiresAt time.Time
	if ttl > 0 {
		expiresAt = time.Now().Add(ttl)
	}

	store.data[key] = Entry{Value: val, ExpiresAt: expiresAt}
}

func (store *Storage) Incr(key string) (int64, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	entry, ok := store.data[key]
	if !ok {
		store.data[key] = Entry{Value: "0"}
		entry = store.data[key]
	}

	intVal, err := strconv.ParseInt(entry.Value, 10, 64)
	if err != nil {
		return 0, errors.New("ERR value is not an integer or out of range")
	}

	intVal++
	store.data[key] = Entry{Value: strconv.FormatInt(intVal, 10)}

	return intVal, nil
}

func (store *Storage) Decr(key string) (int64, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	entry, ok := store.data[key]
	if !ok {
		store.data[key] = Entry{Value: "0"}
		entry = store.data[key]
	}

	intVal, err := strconv.ParseInt(entry.Value, 10, 64)
	if err != nil {
		return 0, errors.New("ERR value is not an integer or out of range")
	}

	intVal--
	store.data[key] = Entry{Value: strconv.FormatInt(intVal, 10)}

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
