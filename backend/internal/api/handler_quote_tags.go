package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/avearmin/wisdomwell/internal/database"
	"github.com/google/uuid"
)

func (c Config) HandlerGetAllQuoteTags(w http.ResponseWriter, r *http.Request) {
	quoteTags, err := c.Db.GetAllQuoteTags(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	outgoing := make([]QuoteTag, len(quoteTags))
	for i, quoteTag := range quoteTags {
		outgoing[i] = dbQuoteTagToJSONQuoteTag(quoteTag)
	}

	if err := respondWithJson(w, http.StatusOK, outgoing); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
	}
}

func (c Config) HandlerGetQuoteTag(w http.ResponseWriter, r *http.Request) {
	quoteIDFromURL := r.URL.Query().Get("quote_id")

	quoteID, err := uuid.Parse(quoteIDFromURL)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "malformed uuid in url")
		return
	}

	tagIDFromURL := r.URL.Query().Get("tag_id")

	tagID, err := uuid.Parse(tagIDFromURL)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "malformed uuid in url")
		return
	}

	quoteTag, err := c.Db.GetQuoteTag(r.Context(), database.GetQuoteTagParams{
		QuoteID: quoteID,
		TagID:   tagID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	outgoing := dbQuoteTagToJSONQuoteTag(quoteTag)
	if err := respondWithJson(w, http.StatusOK, outgoing); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
	}

}
