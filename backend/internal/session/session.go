package session

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	UserID uuid.UUID
	CreatedAt time.Time
	ExpiresAt time.Time
	LastActivity time.Time
}

func (s Session) isExpired() bool {
	return time.Now().After(s.ExpiresAt)
}
