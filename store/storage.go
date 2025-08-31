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
	Type      string
	String    string
	Array     []string
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

	// if ok && !entry.ExpiresAt.IsZero() && time.Now().After(entry.ExpiresAt) {
	// 	store.mu.Lock()
	// 	delete(store.data, key)
	// 	store.mu.Unlock()
	// 	ok = false
	// }

	if !ok {
		return "", false
	}

	return entry.String, true
}

func (store *Storage) Set(key string, val string, expiresAt time.Time) {
	store.mu.Lock()
	defer store.mu.Unlock()

	store.data[key] = Entry{Type: "string", String: val, ExpiresAt: expiresAt}
}

func (store *Storage) LPush(key string, val []string) (int, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	existingData, ok := store.data[key]
	if !ok {
		existingData = Entry{Type: "array", Array: make([]string, 0)}
	}

	if existingData.Type != "array" {
		return 0, errors.New("WRONGTYPE operation against a key holding the wrong kind of value")
	}

	existingData.Array = append(val, existingData.Array...)
	store.data[key] = existingData

	return len(existingData.Array), nil
}

func (store *Storage) RPush(key string, val []string) (int, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	existingData, ok := store.data[key]
	if !ok {
		existingData = Entry{Type: "array", Array: make([]string, 0)}
	}

	if existingData.Type != "array" {
		return 0, errors.New("WRONGTYPE operation against a key holding the wrong kind of value")
	}

	existingData.Array = append(existingData.Array, val...)
	store.data[key] = existingData

	return len(existingData.Array), nil
}

func (store *Storage) LRange(key string, start, stop int64) ([]string, error) {
	store.mu.RLock()
	defer store.mu.RUnlock()

	existingData, ok := store.data[key]
	if !ok {
		existingData = Entry{Type: "array", Array: make([]string, 0)}
	}

	if existingData.Type != "array" {
		return nil, errors.New("WRONGTYPE operation against a key holding the wrong kind of value")
	}

	if start >= int64(len(existingData.Array)) || start > stop {
		return []string{}, nil
	}

	if stop >= int64(len(existingData.Array)) {
		stop = int64(len(existingData.Array))
	}

	if start < 0 && stop < 0 {
		start = int64(len(existingData.Array)) + start
		stop = int64(len(existingData.Array)) + stop
	} else if start < int64(len(existingData.Array))*-1 {
		start = 0
	}

	return existingData.Array[start : stop+1], nil
}

func (store *Storage) Incr(key string) (int64, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	entry, ok := store.data[key]
	if !ok {
		store.data[key] = Entry{Type: "string", String: "0"}
		entry = store.data[key]
	}

	intVal, err := strconv.ParseInt(entry.String, 10, 64)
	if err != nil {
		return 0, errors.New("ERR value is not an integer or out of range")
	}

	intVal++
	store.data[key] = Entry{String: strconv.FormatInt(intVal, 10)}

	return intVal, nil
}

func (store *Storage) Decr(key string) (int64, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	entry, ok := store.data[key]
	if !ok {
		store.data[key] = Entry{Type: "string", String: "0"}
		entry = store.data[key]
	}

	intVal, err := strconv.ParseInt(entry.String, 10, 64)
	if err != nil {
		return 0, errors.New("ERR value is not an integer or out of range")
	}

	intVal--
	store.data[key] = Entry{String: strconv.FormatInt(intVal, 10)}

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
