package api

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"net/http"
)

func (c Config) HandlerGetQuote(w http.ResponseWriter, r *http.Request) {
	incoming := struct {
		ID uuid.UUID
	}{}

	if err := readParameters(r, &incoming); err != nil {
		respondWithError(w, http.StatusBadRequest, "malformed request body")
	}

	quote, err := c.db.GetQuote(r.Context(), incoming.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "internal server error")
		}
	}

	outgoing := dbQuoteToJSONQuote(quote)
	if err := respondWithJson(w, http.StatusOK, outgoing); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
	}
}
