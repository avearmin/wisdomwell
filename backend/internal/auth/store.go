package auth

import (
	"sync"
	"time"
)

type Store struct {
	states  map[string]time.Time
	mu      *sync.RWMutex
	expires time.Duration
}

func NewStore(duration time.Duration) *Store {
	return &Store{
		states:  make(map[string]time.Time),
		mu:      &sync.RWMutex{},
		expires: duration,
	}
}

func (s *Store) AddState(state string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	expiration := time.Now().Add(s.expires)
	s.states[state] = expiration

	time.AfterFunc(s.expires, func() {
		s.DeleteState(state)
	})
}

func (s *Store) ValidateState(state string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	expiration, exists := s.states[state]
	if exists && time.Now().Before(expiration) {
		return true
	}
	return false
}

func (s *Store) DeleteState(state string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.states, state)
}
