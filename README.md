# Gokinde SDK for Go Fiber

The Gokinde SDK is designed to integrate Kinde authentication with Go applications using the Fiber web framework. It handles OAuth2 authentication flow, including login, registration, logout, and user session management.

## Features

- OAuth2 authentication flow
- Session management
- User authentication and route protection
- Easy integration with Fiber applications

## Requirements

- Go 1.15+
- Fiber v2

## Installation

To use the Gokinde SDK in your Go Fiber application, first ensure you have Go and Fiber installed. Then, import the `gokinde` package in your Go application.

## Usage

### Setting up the SDK

1. **Initialization**: Initialize the Gokinde SDK in your main application file where you set up your Fiber app.

   ```go
   package main

   import (
       "github.com/gofiber/fiber/v2"
       "yourmodule/gokinde"
   )

   func main() {
       app := fiber.New()

       gokinde.SetupKinde(app, gokinde.KindeCredentials{
           IssuerBaseUrl:    "https://example.com",
           RedirectUrl:      "https://yourapp.com/redirect",
           SiteUrl:          "https://yourapp.com",
           Secret:           "your-secret",
           UnAuthorisedUrl:  "https://yourapp.com/unauthorized",
           ClientID:         "your-client-id",
       })

       // Define other routes
       // ...

       app.Listen(":3000")
   }
   ```

2. **Configure OAuth2 URLs**: Define the OAuth2 URLs for your application.

   ```go
   func setupKindeURLs() gokinde.KindeURLs {
       return gokinde.KindeURLs{
           SiteUrl:         "http://example.com",
           RedirectUrl:     "http://example.com/redirect",
           UnAuthorisedUrl: "http://example.com/unauthorized",
       }
   }
   ```

### Implementing Handlers

Implement custom handlers for login, registration, logout, and the OAuth2 callback in your Fiber application.

### Session Management

The SDK uses Fiber's session middleware for managing user sessions. Ensure your application handles sessions correctly.

## Contributing

Contributions to the Gokinde SDK are welcome. Please ensure that your code adheres to the existing style and that all tests pass.

## License

Specify your license here or state if it's proprietary.
