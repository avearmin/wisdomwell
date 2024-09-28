package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

func (c Config) HandlerGetAllTags(w http.ResponseWriter, r *http.Request) {
	tags, err := c.db.GetAllTags(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	
	outgoing := make([]Tag, len(tags))
	
	for i, tag := range tags {
		outgoing[i] = dbTagToJSONTag(tag)	
	}

	respondWithJson(w, http.StatusOK, outgoing)
}

func (c Config) HandlerGetTag(w http.ResponseWriter, r *http.Request) {
	idFromURL := r.URL.Query().Get("tag_id")

	id, err := uuid.Parse(idFromURL)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "malformed uuid in url")
		return
	}

	tag, err := c.db.GetTag(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "not found")
			return
		} else {
			respondWithError(w, http.StatusInternalServerError, "internal server error")
			return
		}
	}

	outgoing := dbTagToJSONTag(tag)
	if err := respondWithJson(w, http.StatusOK, outgoing); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}

}
