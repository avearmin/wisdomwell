package api

import "net/http"

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
