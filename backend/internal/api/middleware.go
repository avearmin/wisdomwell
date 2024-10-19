package api

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type authedHandler func(http.ResponseWriter, *http.Request, uuid.UUID)

func (c Config) MiddlewareAuth(next authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		fields := strings.Fields(authHeader)

		if len(fields) < 2 {
			respondWithError(w, http.StatusBadRequest, "malformed authorization header")
		}
		if fields[0] != "sessionID" {
			respondWithError(w, http.StatusBadRequest, "authorization header was not a sessionID")
		}

		sessionID := fields[1]

		session, ok := c.SessionStore.Get(sessionID)
		if !ok {
			respondWithError(w, http.StatusNotFound, "session does not exist")
		}

		userID := session.UserID

		next(w, r, userID)
	}
}
