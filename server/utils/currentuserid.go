package utils

import (
	"strconv"

	"github.com/AndreySpies/doccer/domain/constants"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/labstack/echo/v4"
)

func GetCurrentUserID(context echo.Context, apiKey string) (currentUserID int, err error) {
	cookie, err := context.Cookie(constants.JWTCookieName)
	if err != nil {
		return currentUserID, err
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(apiKey), nil
	})
	if err != nil {
		return currentUserID, err
	}

	claims := token.Claims.(*jwt.StandardClaims)
	currentUserID, err = strconv.Atoi(claims.Issuer)

	return currentUserID, err
}
