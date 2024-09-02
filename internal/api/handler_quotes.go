package api

import (
	"database/sql"
	"errors"
	"github.com/avearmin/wisdomwell/internal/database"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (c Config) HandlerGetQuote(w http.ResponseWriter, r *http.Request) {
	incoming := struct {
		ID uuid.UUID `json:"id"`
	}{}

	if err := readParameters(r, &incoming); err != nil {
		respondWithError(w, http.StatusBadRequest, "malformed request body")
		return
	}

	quote, err := c.db.GetQuote(r.Context(), incoming.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "not found")
			return
		} else {
			respondWithError(w, http.StatusInternalServerError, "internal server error")
			return
		}
	}

	outgoing := dbQuoteToJSONQuote(quote)
	if err := respondWithJson(w, http.StatusOK, outgoing); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}
}

func (c Config) HandlerPostQuote(w http.ResponseWriter, r *http.Request) {
	incoming := struct {
		Content string    `json:"content"`
		UserID  uuid.UUID `json:"user_id"`
	}{}

	if err := readParameters(r, &incoming); err != nil {
		respondWithError(w, http.StatusBadRequest, "malformed request body")
		return
	}

	quote, err := c.db.PostQuote(r.Context(), database.PostQuoteParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Content:   incoming.Content,
		UserID:    incoming.UserID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	outgoing := dbQuoteToJSONQuote(quote)
	if err := respondWithJson(w, http.StatusOK, outgoing); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
	}
}
