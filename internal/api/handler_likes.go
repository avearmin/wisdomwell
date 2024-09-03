package api

import (
	"database/sql"
	"errors"
	"github.com/avearmin/wisdomwell/internal/database"
	"github.com/google/uuid"
	"net/http"
)

func (c Config) HandlerGetLike(w http.ResponseWriter, r *http.Request) {
	incoming := struct {
		UserID  uuid.UUID `json:"user_id"`
		QuoteID uuid.UUID `json:"quote_id"`
	}{}

	if err := readParameters(r, &incoming); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	like, err := c.db.GetLike(r.Context(), database.GetLikeParams{
		UserID:  incoming.UserID,
		QuoteID: incoming.QuoteID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "not found")
			return
		} else {
			respondWithError(w, http.StatusInternalServerError, "internal server error")
			return
		}
	}

	outgoing := dbLikeToJSONLike(like)
	if err := respondWithJson(w, http.StatusOK, outgoing); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
	}
}
