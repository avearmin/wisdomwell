package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token="

type GoogleUserData struct {
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
}

func GenerateGoogleURL(clientID, clientSecret, redirectURL, state string) string {
	conf := generateGoogleOauth2Config(clientID, clientSecret, redirectURL)
	url := conf.AuthCodeURL(state)
	return url
}

func GetGoogleUserData(clientID, clientSecret, redirectURL, code string, ctx context.Context) (GoogleUserData, error) {
	conf := generateGoogleOauth2Config(clientID, clientSecret, redirectURL)

	token, err := conf.Exchange(ctx, code)
	if err != nil {
		return GoogleUserData{}, err
	}

	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return GoogleUserData{}, err
	}

	decoder := json.NewDecoder(response.Body)

	var userData GoogleUserData
	if err := decoder.Decode(&userData); err != nil {
		return GoogleUserData{}, err
	}

	if !userData.VerifiedEmail {
		return GoogleUserData{}, errors.New("email not verified")
	}

	return userData, nil
}

func generateGoogleOauth2Config(clientID, clientSecret, redirectURL string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}
