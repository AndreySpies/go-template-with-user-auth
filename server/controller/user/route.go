package user

import (
	"database/sql"

	"github.com/AndreySpies/doccer/data/repository"
	userservice "github.com/AndreySpies/doccer/domain/service/user"
	"github.com/labstack/echo/v4"
)

func CreateUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		controller := getUserController(db)
		return controller.createUser(c)
	}
}

func Login(db *sql.DB, apiKey string) echo.HandlerFunc {
	return func(c echo.Context) error {
		controller := getUserController(db)
		return controller.login(c, apiKey)
	}
}

func Logout(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		controller := getUserController(db)
		return controller.logout(c)
	}
}

func getUserController(db *sql.DB) *userController {
	repo := repository.NewUserRepo(db)
	service := userservice.NewUserService(repo)
	controller := newUserController(service)

	return controller
}
