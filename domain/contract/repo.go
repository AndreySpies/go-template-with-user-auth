package contract

import (
	"github.com/AndreySpies/doccer/domain/entity"
	"github.com/AndreySpies/doccer/server/request"
)

type UserRepo interface {
	Create(user *request.CreateUser, encryptedPassword []byte) (err error)
	FindOneByEmail(email string) (user entity.User, err error)
	FindOneByID(id int) (user entity.User, err error)
}
