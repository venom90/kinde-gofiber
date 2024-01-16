package gokinde

import (
	"golang.org/x/oauth2"
)

func NewOAuthConfig(clientID, clientSecret, redirectURL, authURL, tokenURL string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL,
			TokenURL: tokenURL,
		},
		Scopes: []string{"openid", "profile", "email"},
	}
}
