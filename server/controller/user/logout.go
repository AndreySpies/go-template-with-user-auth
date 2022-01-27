package user

import (
	"net/http"

	"github.com/AndreySpies/doccer/domain/constants"
	"github.com/labstack/echo/v4"
)

func (c *userController) logout(ctx echo.Context) error {
	cookie := c.userService.Logout()
	ctx.SetCookie(cookie)

	return ctx.JSON(http.StatusOK, constants.EmptyString)
}
