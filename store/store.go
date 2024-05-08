package store

import (
	"sync"
)

/*
Store the values in a simple key-values map and
use a mutex to make sure that the data is accessed
because there is multiple goroutines that is using this store.
*/
type Store struct {
	data map[string]string
	mu   sync.RWMutex
}

func New() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

// Sets function will add or update the values in the store.
func (s *Store) Set(key, value string) {
	s.mu.Lock()  // Lock for writing
	defer s.mu.Unlock()	
	s.data[key] = value
}

// Gets function will retrieve the values from the store.
// Returns the values and boolean (the key is found or not)
func (s *Store) Get(key string) (string,bool) {
	s.mu.RLock()  // Lock for reading
	defer s.mu.RUnlock()
	value, ok := s.data[key]
	return value, ok
}