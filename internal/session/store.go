package session

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type Store struct {
	sessions map[string]*Session
	duration time.Duration
	mu *sync.RWMutex
}

func NewStore(duration time.Duration) Store {
	return Store{
		sessions: make(map[string]*Session),
		duration: duration,
		mu: &sync.RWMutex{},
	}
}

func (s Store) Get(sessionID string) (*Session, bool) {
	session, ok := s.sessions[sessionID]
	if !ok {
		return nil, false
	}
	return session, true
}

func (s Store) CreateSession(userID uuid.UUID) string {
	timeNow := time.Now()
	session := &Session{
		UserID: userID,
		CreatedAt: timeNow,
		ExpiresAt: timeNow.Add(s.duration),
		LastActivity: timeNow,
	}

	sessionID := uuid.New().String()
	s.addSession(sessionID, session)
	time.AfterFunc(s.duration, func() {
		s.deleteSession(sessionID)
	})

	return sessionID
}

func (s Store) HasSession(id string) bool {
	if _, ok := s.sessions[id]; !ok {
		return false
	}
	return true
}

func (s Store) addSession(id string, session *Session) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.sessions[id] = session
}

func (s Store) deleteSession(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.sessions, id)
}
