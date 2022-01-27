package user

import "github.com/AndreySpies/doccer/domain/contract"

type userController struct {
	userService contract.UserService
}

func newUserController(s contract.UserService) *userController {
	return &userController{
		userService: s,
	}
}
