package contract

import (
	"net/http"

	"github.com/AndreySpies/doccer/server/request"
)

type UserService interface {
	CreateUser(user *request.CreateUser) (err error)
	Login(userData *request.Login, apiKey string) (cookie *http.Cookie, authenticated bool, err error)
	Logout() (cookie *http.Cookie)
}

type DashboardService interface {
}
