package main

import (
	"fmt"

	"github.com/AndreySpies/doccer/infra/config"
	"github.com/AndreySpies/doccer/infra/database"
	"github.com/AndreySpies/doccer/server/middleware"
	"github.com/AndreySpies/doccer/server/route"
	"github.com/labstack/echo/v4"
)

// AppVersion injected by linker flags
var AppVersion string

func main() {
	e := echo.New()

	cfg, err := config.Read()
	if err != nil {
		e.Logger.Fatal(err)
		panic(err)
	}

	cfg.Version = AppVersion

	db, err := database.NewMySQLDB(
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
	if db == nil || err != nil {
		e.Logger.Fatal(err)
		panic(err)
	}

	middleware.SetupMiddlewares(e, cfg)
	route.SetupRoutes(e, cfg, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Port)))
}
