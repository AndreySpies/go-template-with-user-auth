package dashboard

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/AndreySpies/doccer/domain/constants"
	"github.com/AndreySpies/doccer/server/utils"
	"github.com/labstack/echo/v4"
)

func Dashboard(db *sql.DB, apiKey string) echo.HandlerFunc {
	return func(c echo.Context) error {
		currentUser, err := utils.GetCurrentUser(c, db, apiKey)
		if err != nil {
			return c.JSON(
				http.StatusUnauthorized,
				constants.ErrorMessage(constants.SomethingWentWrongErrorCode),
			)
		}

		return c.JSON(
			http.StatusOK,
			fmt.Sprintf(
				constants.GreetingUser,
				currentUser.FirstName,
			),
		)
	}
}
