package main

import (
	// "fmt"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Dummy user for example
var user = struct {
	Email    string
	Password string
}{
	Email:    "user@example.com",
	Password: "password123",
}

func login(secretKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type LoginRequest struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		var request LoginRequest
		if err := c.BodyParser(&request); err != nil {
			return err
		}

		// Check credentials - In real world, you should check against a database
		if request.Email != user.Email || request.Password != user.Password {
			return fiber.ErrUnauthorized
		}

		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["email"] = user.Email
		claims["role"] = "admin" // example role
		claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

		// Generate encoded token
		t, err := token.SignedString([]byte(secretKey))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{"token": t})
	}
}

// func isAdmin(c *fiber.Ctx) error {
// 	// Extract the token from the Fiber context (inserted by the JWT middleware)
// 	token := c.Locals("user").(*jwt.Token)
// 	claims := token.Claims.(jwt.MapClaims)

// 	if claims["role"] != "admin" {
// 		return fiber.ErrUnauthorized
// 	}
// 	return c.Next()
// }

// const userContextKey = "user"

type UserData struct {
	Email string
	Role  string
}

func isAdmin(c *fiber.Ctx) error {
	// Safely assert the type of token from c.Locals
	token, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return fiber.ErrUnauthorized
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fiber.ErrUnauthorized
	}

	// Check for the admin role in the claims
	if claims["role"] != "admin" {
		return fiber.ErrUnauthorized
	}

	return c.Next()
}

// // extractUserFromJWT is a middleware that extracts user data from the JWT token
// func extractUserFromJWT(c *fiber.Ctx) error {
// 	user := &UserData{}

// 	// Extract the token from the Fiber context (inserted by the JWT middleware)
// 	token := c.Locals("user").(*jwt.Token)
// 	claims := token.Claims.(jwt.MapClaims)

// 	fmt.Println(claims)

// 	user.Email = claims["email"].(string)
// 	user.Role = claims["role"].(string)

// 	// Store the user data in the Fiber context
// 	c.Locals(userContextKey, user)

// 	return c.Next()
// }

func extractUserFromJWT(c *fiber.Ctx) error {
	token := c.Locals("user") // JWT middleware should have set this
	if token == nil {
		return fiber.ErrUnauthorized
	}

	// Ensure it's a *jwt.Token
	jwtToken, ok := token.(*jwt.Token)
	if !ok {
		return fiber.ErrUnauthorized
	}

	// Optionally, you could store user data in locals for later use
	c.Locals("jwt_token", jwtToken)
	return c.Next()
}
