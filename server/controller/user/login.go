package user

import (
	"net/http"

	"github.com/AndreySpies/doccer/domain/constants"
	"github.com/AndreySpies/doccer/server/request"
	"github.com/AndreySpies/doccer/server/response"
	"github.com/labstack/echo/v4"
)

func (c *userController) login(ctx echo.Context, apiKey string) error {
	userData := &request.Login{}
	err := ctx.Bind(userData)
	if err != nil {
		return err
	}

	cookie, authenticated, err := c.userService.Login(userData, apiKey)
	if err != nil {
		return err
	}

	if !authenticated {
		result := response.Login{Message: constants.UserLoginInvalid}
		return ctx.JSON(http.StatusBadRequest, result)
	}

	ctx.SetCookie(cookie)

	return ctx.JSON(http.StatusOK, constants.EmptyString)
}
