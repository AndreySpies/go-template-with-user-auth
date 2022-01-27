package utils

import (
	"database/sql"

	"github.com/AndreySpies/doccer/data/repository"
	"github.com/AndreySpies/doccer/domain/entity"
	"github.com/labstack/echo/v4"
)

func GetCurrentUser(context echo.Context, db *sql.DB, apiKey string) (currentUser entity.User, err error) {
	currentUserID, err := GetCurrentUserID(context, apiKey)
	if err != nil {
		return currentUser, err
	}

	repo := repository.NewUserRepo(db)
	currentUser, err = repo.FindOneByID(currentUserID)

	return currentUser, err
}
