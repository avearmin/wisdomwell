package session

import (
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Store struct {
	sessions map[string]*Session
	duration time.Duration
	cookieName string
	mu *sync.RWMutex
}

func NewStore(duration time.Duration, cookieName string) Store {
	return Store{
		sessions: make(map[string]*Session),
		duration: duration,
		cookieName: cookieName,
		mu: &sync.RWMutex{},
	}
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
