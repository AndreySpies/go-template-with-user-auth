package userservice

import (
	"time"

	"github.com/AndreySpies/doccer/server/request"
	"golang.org/x/crypto/bcrypt"
)

func (s *userService) CreateUser(user *request.CreateUser) (err error) {
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	timeLayout := "2006-01-02"
	user.CreatedAt = time.Now().Format(timeLayout)

	err = s.userRepo.Create(user, encryptedPassword)

	return err
}
