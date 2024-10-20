package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/avearmin/wisdomwell/internal/database"
	"github.com/google/uuid"
)

func (c Config) HandlerGetAllLikes(w http.ResponseWriter, r *http.Request) {
	likes, err := c.Db.GetAllLikes(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	outgoing := make([]Like, len(likes))
	for i, like := range likes {
		outgoing[i] = dbLikeToJSONLike(like)
	}

	if err := respondWithJson(w, http.StatusOK, outgoing); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
	}
}

func (c Config) HandlerGetLike(w http.ResponseWriter, r *http.Request) {
	quoteIDFromURL := r.URL.Query().Get("quote_id")

	quoteID, err := uuid.Parse(quoteIDFromURL)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "malformed uuid in url")
		return
	}

	userIDFromURL := r.URL.Query().Get("user_id")

	userID, err := uuid.Parse(userIDFromURL)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "malformed uuid in url")
		return
	}

	like, err := c.Db.GetLike(r.Context(), database.GetLikeParams{
		UserID:  userID,
		QuoteID: quoteID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	outgoing := dbLikeToJSONLike(like)
	if err := respondWithJson(w, http.StatusOK, outgoing); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
	}
}

func (c Config) HandlerPostLike(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	incoming := struct {
		QuoteID uuid.UUID `json:"quote_id"`
	}{}

	if err := readParameters(r, &incoming); err != nil {
		respondWithError(w, http.StatusBadRequest, "malformed request body")
		return
	}

	like, err := c.Db.PostLike(r.Context(), database.PostLikeParams{
		UserID:    userID,
		QuoteID:   incoming.QuoteID,
		CreatedAt: time.Now(),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	outgoing := dbLikeToJSONLike(like)
	if err := respondWithJson(w, http.StatusOK, outgoing); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
	}
}

func (c Config) HandlerDeleteLike(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	incoming := struct {
		QuoteID uuid.UUID `json:"quote_id"`
	}{}

	if err := readParameters(r, &incoming); err != nil {
		respondWithError(w, http.StatusBadRequest, "malformed request body")
		return
	}

	err := c.Db.DeleteLike(r.Context(), database.DeleteLikeParams{
		UserID:  userID,
		QuoteID: incoming.QuoteID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "internal server error")
		}
		return
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
