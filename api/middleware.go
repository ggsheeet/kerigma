package api

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	environment := os.Getenv("ENVIRONMENT")
	var allowedOrigin string
	token := os.Getenv("AUTH_TOKEN")

	if environment == "development" || environment == "docker" {
		allowedOrigin = os.Getenv("AUTH_ORIGIN_DEV")
	} else if environment == "production" {
		allowedOrigin = os.Getenv("AUTH_ORIGIN_PRDO")
	}

	return func(c echo.Context) error {
		origin := c.Request().Header.Get("Origin")

		// Check if the origin is allowed
		if origin != allowedOrigin {
			return echo.NewHTTPError(http.StatusForbidden, "Forbidden: Invalid Origin")
		}

		// CORS headers
		c.Response().Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		// Handle preflight requests
		if c.Request().Method == http.MethodOptions {
			return c.NoContent(http.StatusNoContent)
		}

		// Token-based authentication
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader != "Bearer "+token {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized: Invalid Token")
		}

		return next(c)
	}
}