package middleware

import (
	"github.com/AndreySpies/doccer/infra/config"
	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"
)

func SetupMiddlewares(e *echo.Echo, cfg *config.Config) {
	e.Use(
		NewLoggerMiddleware(),
		echo_middleware.Recover(),
		echo_middleware.AddTrailingSlash(),
		echo_middleware.RequestID(),
		echo_middleware.CORS(),
		echo_middleware.Gzip(),
		echo_middleware.Secure(),
		echo_middleware.BodyLimit(cfg.HTTPBodySizeLimit),
	)
}
