package api

import (
	"database/sql"
	"errors"
	"github.com/avearmin/wisdomwell/internal/database"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (c Config) HandlerGetAllQuotes(w http.ResponseWriter, r *http.Request) {
	quotes, err := c.db.GetAllQuotes(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	outgoing := make([]Quote, len(quotes))
	for i, quote := range quotes {
		outgoing[i] = dbQuoteToJSONQuote(quote)
	}

	if err := respondWithJson(w, http.StatusOK, outgoing); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
	}
}

func (c Config) HandlerGetQuote(w http.ResponseWriter, r *http.Request) {
	idFromURL := r.URL.Query().Get("quote_id")

	id, err := uuid.Parse(idFromURL)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "malformed uuid in url")
		return
	}

	quote, err := c.db.GetQuote(r.Context(), id)
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

func (c Config) HandlerPostQuote(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	incoming := struct {
		Content string `json:"content"`
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
		UserID:    userID,
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

func (c Config) HandlerDeleteQuote(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
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
		} else {
			respondWithError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}
	if quote.UserID != userID {
		respondWithError(w, http.StatusForbidden, "forbidden")
		return
	}

	if err := c.db.DeleteQuote(r.Context(), incoming.ID); err != nil {
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

func (c Config) HandlerGetRandomQuote(w http.ResponseWriter, r *http.Request) {
	quote, err := c.db.GetRandomQuote(r.Context())
	if err != nil {	
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "not found")		
		} else {
			respondWithError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	outgoing := dbQuoteToJSONQuote(quote)
	if err := respondWithJson(w, http.StatusOK, outgoing); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
	}
}
