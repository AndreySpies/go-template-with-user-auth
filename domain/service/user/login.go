package userservice

import (
	"net/http"
	"strconv"
	"time"

	"github.com/AndreySpies/doccer/domain/constants"
	"github.com/AndreySpies/doccer/domain/entity"
	"github.com/AndreySpies/doccer/server/request"
	"github.com/dgrijalva/jwt-go/v4"
	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Login(userData *request.Login, apiKey string) (cookie *http.Cookie, authenticated bool, err error) {
	authenticated = false

	user, validated, err := findAndValidateUser(s, userData)
	if err != nil || !validated {
		return cookie, authenticated, err
	}

	claims := buildClaims(user)

	token, err := claims.SignedString([]byte(apiKey))
	if err != nil {
		return cookie, authenticated, err
	}

	cookie = buildCookie(token)

	authenticated = true

	return cookie, authenticated, err
}

func findAndValidateUser(s *userService, userData *request.Login) (user entity.User, validated bool, err error) {
	validated = false

	user, err = s.userRepo.FindOneByEmail(userData.Email)
	if err != nil || user.Id == 0 {
		return user, validated, err
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(userData.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			err = nil
		}
		return user, validated, err
	}

	validated = true
	return user, validated, err
}

func buildClaims(user entity.User) *jwt.Token {
	expiration := float64(time.Now().Add(time.Hour * 24).Unix())

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(user.Id),
		ExpiresAt: jwt.NewTime(expiration),
	})

	return claims
}

func buildCookie(token string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = constants.JWTCookieName
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Hour * 24)
	cookie.HttpOnly = true

	return cookie
}
