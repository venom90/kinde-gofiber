package gokinde

import (
	"log"
	"strings"

	"github.com/MicahParks/keyfunc"
	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
)

func JWTVerify(issuer string, audience string) fiber.Handler {
	jwksURL := issuer + "/.well-known/jwks.json"

	// Creating options for keyfunc.Get
	options := keyfunc.Options{
		RefreshInterval: 0, // Default value, no automatic refresh
		RefreshErrorHandler: func(err error) {
			log.Printf("There was an error with the jwt.Keyfunc: %s", err.Error())
		},
		RefreshUnknownKID: true,
	}

	jwks, err := keyfunc.Get(jwksURL, options)
	if err != nil {
		log.Fatalf("Failed to get JWKS from the given URL: %v", err)
	}

	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse and verify the token
		token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwks.Keyfunc(token)
		})
		if err != nil {
			log.Printf("Token not valid: %v\n", err)
			return c.SendStatus(fiber.StatusForbidden)
		}

		if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
			if audience != "" && !claims.VerifyAudience(audience, true) {
				log.Println("Invalid audience in token")
				return c.SendStatus(fiber.StatusForbidden)
			}

			if !claims.VerifyIssuer(issuer, true) {
				log.Println("Invalid issuer in token")
				return c.SendStatus(fiber.StatusForbidden)
			}

			log.Println("Token is valid")
			return c.Next()
		}

		log.Println("Invalid token claims")
		return c.SendStatus(fiber.StatusForbidden)
	}
}
