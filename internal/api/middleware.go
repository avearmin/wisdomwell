package api

import (
	"net/http"
	"strings"
)

func (c Config) MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		fields := strings.Fields(authHeader)

		if len(fields) < 2 {
			respondWithError(w, http.StatusBadRequest, "malformed authorization header")
		}
		if fields[0] != "sessionID" {
			respondWithError(w, http.StatusBadRequest, "authorization header was not a sessionID")
		}

		sessionID := fields[1]

		if ok := c.sessionStore.HasSession(sessionID); !ok {
			respondWithError(w, http.StatusNotFound, "session does not exist")
		}

		next.ServeHTTP(w, r)
	})
}
