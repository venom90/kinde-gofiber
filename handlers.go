package gokinde

import (
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

func defineRoutes(app *fiber.App, urls KindeURLs) {
	app.Get("/login", loginHandler)
	app.Get("/register", registerHandler)
	app.Get("/logout", logoutHandler)
	app.Get("/kinde_callback", kindeCallbackHandler(urls))
}

func loginHandler(c *fiber.Ctx) error {
	state, err := RandomString(32)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error generating state")
	}

	loginUrl := constructAuthURL(state, "")
	err = storeSessionValue(c, "kindeState", state)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error saving state in session")
	}

	return c.Redirect(loginUrl)
}

func registerHandler(c *fiber.Ctx) error {
	state, err := RandomString(32)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error generating state")
	}

	registerUrl := constructAuthURL(state, "registration")
	err = storeSessionValue(c, "kindeState", state)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error saving state in session")
	}

	return c.Redirect(registerUrl)
}

func logoutHandler(c *fiber.Ctx) error {
	// Destroy the session
	err := destroySession(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error destroying session")
	}

	logoutUrl := fmt.Sprintf("%s/logout?redirect=%s", issuerUrl, url.QueryEscape(issuerUrl))
	return c.Redirect(logoutUrl)
}

func kindeCallbackHandler(urls KindeURLs) fiber.Handler {
	return func(c *fiber.Ctx) error {
		session, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		state, _ := session.Get("kindeState").(string)
		queryState := c.Query("state")

		if state != queryState {
			return c.Redirect(urls.UnAuthorisedUrl)
		}

		code := c.Query("code")
		token, err := client.Exchange(c.Context(), code, oauth2.SetAuthURLParam("redirect_uri", urls.SiteUrl+"/kinde_callback"))
		if err != nil {
			// Log the error and redirect to unauthorized URL or default redirect URL
			return c.Redirect(urls.UnAuthorisedUrl)
		}

		// Store the token in the session
		session.Set("kindeAccessToken", token.AccessToken)
		session.Delete("kindeState")
		err = session.Save()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to save session: " + err.Error())
		}

		return c.Redirect(urls.RedirectUrl)
	}
}

func constructAuthURL(state string, page string) string {
	authURL := client.AuthCodeURL(state)
	if page != "" {
		authURL += "&start_page=" + page
	}
	return authURL
}

func storeSessionValue(c *fiber.Ctx, key string, value interface{}) error {
	sess, err := store.Get(c)
	if err != nil {
		return err
	}

	sess.Set(key, value)
	return sess.Save()
}

func destroySession(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return err
	}

	return sess.Destroy()
}
