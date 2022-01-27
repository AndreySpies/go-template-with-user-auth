package userservice

import "github.com/AndreySpies/doccer/domain/contract"

type userService struct {
	userRepo contract.UserRepo
}

func NewUserService(repo contract.UserRepo) contract.UserService {
	return &userService{
		userRepo: repo,
	}
}
