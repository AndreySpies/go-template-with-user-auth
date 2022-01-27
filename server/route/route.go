package route

import (
	"database/sql"

	"github.com/AndreySpies/doccer/infra/config"
	"github.com/AndreySpies/doccer/server/controller/dashboard"
	"github.com/AndreySpies/doccer/server/controller/healthcheck"
	"github.com/AndreySpies/doccer/server/controller/user"
	"github.com/AndreySpies/doccer/server/controller/version"
	"github.com/AndreySpies/doccer/server/middleware"
	"github.com/labstack/echo/v4"
	echo_middleware "github.com/labstack/echo/v4/middleware"
)

// SetupRoutes configures API URIs and controllers
func SetupRoutes(e *echo.Echo, cfg *config.Config, db *sql.DB) {
	publicGroup := e.Group("")
	apiKeyAuthGroup := e.Group("", middleware.NewAuthMiddleware(cfg.ApiKey))

	jwtConfig := echo_middleware.JWTConfig{
		SigningKey: []byte(cfg.ApiKey),
	}

	authUserGroup := e.Group("", echo_middleware.JWTWithConfig(jwtConfig))

	setupDefaultRoutes(publicGroup, cfg)
	setupApiKeyAuthRoutes(apiKeyAuthGroup, cfg, db)
	setupAuthUserRoutes(authUserGroup, cfg, db)
}

func setupDefaultRoutes(publicGroup *echo.Group, cfg *config.Config) {
	publicGroup.GET("/health", healthcheck.NewHealthcheckRoute())
	publicGroup.GET("/version", version.NewVersionRoute(cfg.Version))
}

func setupApiKeyAuthRoutes(apiKeyAuthGroup *echo.Group, cfg *config.Config, db *sql.DB) {
	apiKeyAuthGroup.POST("/create_user", user.CreateUser(db))
	apiKeyAuthGroup.GET("/login", user.Login(db, cfg.ApiKey))
}

func setupAuthUserRoutes(authUserGroup *echo.Group, cfg *config.Config, db *sql.DB) {
	authUserGroup.GET("/dashboard", dashboard.Dashboard(db, cfg.ApiKey))
	authUserGroup.POST("/logout", user.Logout(db))
}
