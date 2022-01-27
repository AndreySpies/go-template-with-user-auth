package middleware

import (
	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"
)

func NewAuthMiddleware(apiKey string) echo.MiddlewareFunc {
	return echo_middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == apiKey, nil
	})
}
