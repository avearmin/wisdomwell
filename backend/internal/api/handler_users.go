package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/avearmin/wisdomwell/internal/database"
	"github.com/google/uuid"
)

func (c Config) HandlerGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.db.GetAllUsers(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	outgoing := make([]User, len(users))
	for i, user := range users {
		outgoing[i] = dbUserToJSONUser(user)
	}

	if err := respondWithJson(w, http.StatusOK, outgoing); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
	}
}

func (c Config) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	incoming := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{}

	if err := readParameters(r, &incoming); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	user, err := c.db.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     incoming.Email,
		Name:      incoming.Name,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	outgoing := dbUserToJSONUser(user)
	if err := respondWithJson(w, http.StatusOK, outgoing); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}
}

func (c Config) HandlerGetUser(w http.ResponseWriter, r *http.Request) {
	idFromURL := r.URL.Query().Get("user_id")

	id, err := uuid.Parse(idFromURL)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "malformed uuid in url")
		return
	}

	user, err := c.db.GetUser(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	outgoing := dbUserToJSONUser(user)
	if err := respondWithJson(w, http.StatusOK, outgoing); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}
}

func (c Config) HandlerGetAllQuotesFromUser(w http.ResponseWriter, r *http.Request) {
	idFromURL := r.URL.Query().Get("user_id")

	id, err := uuid.Parse(idFromURL)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "malformed uuid in url")
		return
	}

	quotes, err := c.db.GetAllQuotesFromUser(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "internal server error")
		}
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

func (c Config) HandlerDeleteUser(w http.ResponseWriter, r *http.Request, userID uuid.UUID) {
	if err := c.db.DeleteUser(r.Context(), userID); err != nil {
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
