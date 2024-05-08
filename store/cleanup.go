package store

import "time"

func (s *Store) CleanUpExpiredKey() {
	// Cleanup Interval
	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		for range ticker.C {
			s.removeExpiredKeys()
		}
	}()
}

func (s *Store) removeExpiredKeys() {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	for key, entry := range s.data {
		if !entry.expiryTime.IsZero() && entry.expiryTime.Before(now) {
			delete(s.data, key)
		}
	}
}
