package store

import (
	"sync"
	"time"
)

/*
Store the values in a simple key-values map and
use a mutex to make sure that the data is accessed
because there is multiple goroutines that is using this store.
*/
type Store struct {
	data map[string]*valueEntry
	mu   sync.RWMutex
}

type valueEntry struct {
	value      string
	expiryTime time.Time // 0 means there is no expiry time
}

func New() *Store {
	return &Store{
		data: make(map[string]*valueEntry),
	}
}

// Sets function will add or update the values in the store.
func (s *Store) Set(key, value string, expiryMillis int64) {
	s.mu.Lock() // Lock for writing
	defer s.mu.Unlock()


	// Define the value entry and theres expiry time
	entry := &valueEntry{
		value: value,
	}
	if expiryMillis > 0 {
		entry.expiryTime = time.Now().Add(time.Duration(expiryMillis) * time.Millisecond)
	}

	s.data[key] = entry
}

// Gets function will retrieve the values from the store.
// Returns the values and boolean (the key is found or not)
func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock() // Lock for reading
	defer s.mu.RUnlock()

	entry, ok := s.data[key]
	if !ok {
		return "", false
	}

	// Check if the key has expired
	if !entry.expiryTime.IsZero() && entry.expiryTime.Before(time.Now()) {
		return "", false
	}

	return entry.value, true
}
