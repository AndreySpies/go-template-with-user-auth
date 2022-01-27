package userservice

import (
	"net/http"
	"time"

	"github.com/AndreySpies/doccer/domain/constants"
)

func (s *userService) Logout() (cookie *http.Cookie) {
	return buildLogoutCookie()
}

func buildLogoutCookie() *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = constants.JWTCookieName
	cookie.Value = constants.EmptyString
	cookie.Expires = time.Now().Add(-time.Hour)
	cookie.HttpOnly = true

	return cookie
}
