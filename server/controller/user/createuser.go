package user

import (
	"net/http"

	"github.com/AndreySpies/doccer/domain/constants"
	"github.com/AndreySpies/doccer/server/request"
	"github.com/AndreySpies/doccer/server/response"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

func (c *userController) createUser(ctx echo.Context) error {
	user := &request.CreateUser{}
	err := ctx.Bind(user)
	if err != nil {
		return err
	}

	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		return err
	}

	result := response.CreateUser{}

	err = c.userService.CreateUser(user)
	if err != nil {
		emailAlreadyUsedError := constants.CustomError(constants.EmailAlreadyUsedErrorCode).Error()
		if err.Error() == emailAlreadyUsedError {
			result.Message = constants.ErrorMessage(constants.EmailAlreadyUsedErrorCode)
			return ctx.JSON(http.StatusUnprocessableEntity, result)
		}

		return err
	}

	result.Message = constants.UserCreated
	return ctx.JSON(http.StatusOK, result)
}
