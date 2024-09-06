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

func (c Config) HandlerPostLike(w http.ResponseWriter, r *http.Request) {
	incoming := struct {
		UserID  uuid.UUID `json:"user_id"`
		QuoteID uuid.UUID `json:"quote_id"`
	}{}

	if err := readParameters(r, &incoming); err != nil {
		respondWithError(w, http.StatusBadRequest, "malformed request body")
		return
	}

	like, err := c.db.PostLike(r.Context(), database.PostLikeParams{
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

func (c Config) HandlerDeleteLike(w http.ResponseWriter, r *http.Request) {
	incoming := struct {
		UserID uuid.UUID `json:"user_id"`
		QuoteID uuid.UUID `json:"quote_id"`
	}{}

	if err := readParameters(r, &incoming); err != nil {
		respondWithError(w, http.StatusBadRequest, "malformed request body")
	}

	err := c.db.DeleteLike(r.Context(), database.DeleteLikeParams{
		UserID: incoming.UserID,
		QuoteID: incoming.QuoteID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "internal server error")
		}
	}
	outgoing := struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	}
	if err := respondWithJson(w, http.StatusOK, outgoing); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
	}

}
