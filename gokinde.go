package gokinde

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/oauth2"
)

var (
	store     *session.Store
	client    *oauth2.Config
	issuerUrl string
)

func SetupKinde(app *fiber.App, credentials KindeCredentials, urls KindeURLs) {
	store = session.New(session.Config{
		Expiration: 24 * 60 * 60, // 24 hours
	})

	client = &oauth2.Config{
		ClientID:     credentials.ClientID,
		ClientSecret: credentials.Secret,
		RedirectURL:  credentials.RedirectUrl,
		Endpoint: oauth2.Endpoint{
			AuthURL:  credentials.IssuerBaseUrl + "/oauth2/auth",
			TokenURL: credentials.IssuerBaseUrl + "/oauth2/token",
		},
		Scopes: []string{"openid", "profile", "email"},
	}

	issuerUrl = credentials.IssuerBaseUrl
	// unAuthorisedRedirectUrl = credentials.UnAuthorisedUrl

	defineRoutes(app, urls)
}
