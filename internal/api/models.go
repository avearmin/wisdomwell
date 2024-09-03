package api

import (
	"github.com/avearmin/wisdomwell/internal/database"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
}

func dbUserToJSONUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
		Name:      user.Name,
	}
}

type Quote struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
	UserID    uuid.UUID `json:"user_id"`
	Content   string    `json:"content"`
}

func dbQuoteToJSONQuote(quote database.Quote) Quote {
	return Quote{
		ID:        quote.ID,
		CreatedAt: quote.CreatedAt,
		UpdatedAt: quote.UpdatedAt,
		UserID:    quote.UserID,
		Content:   quote.Content,
	}
}

type Like struct {
	UserID  uuid.UUID `json:"user_id"`
	QuoteID uuid.UUID `json:"quote_id"`
}

func dbLikeToJSONLike(like database.Like) Like {
	return Like{
		UserID:  like.UserID,
		QuoteID: like.QuoteID,
	}
}
