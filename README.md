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

```go
go get github.com/venom90/kinde-gofiber
```

### Setting up the SDK

1. **Initialization**: Initialize the Gokinde SDK in your main application file where you set up your Fiber app.

   ```go
   package main

   import (
       "github.com/gofiber/fiber/v2"
       gokinde "github.com/venom90/kinde-gofiber"
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

### Implementing Handlers

Implement custom handlers for login, registration, logout, and the OAuth2 callback in your Fiber application.

### Session Management

The SDK uses Fiber's session middleware for managing user sessions. Ensure your application handles sessions correctly.

## Contributing

Contributions to the Gokinde SDK are welcome. Please ensure that your code adheres to the existing style and that all tests pass.

## License

MIT License

Copyright (c) 2024 venom90

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
