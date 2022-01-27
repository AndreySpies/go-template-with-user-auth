package version

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewVersionRoute(apiVersion string) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, apiVersion)
	}
}
