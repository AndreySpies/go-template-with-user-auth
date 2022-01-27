package healthcheck

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewHealthcheckRoute() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]bool{
			"ok": true,
		})
	}
}
