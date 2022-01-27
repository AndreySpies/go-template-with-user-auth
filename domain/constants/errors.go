package constants

import "fmt"

const (
	MySQLDuplicateEntryErrorCode = 1062
)

const (
	SomethingWentWrongErrorCode = 2001
	EmailAlreadyUsedErrorCode   = 2002
)

var errorCodeToMessage = map[int]string{
	2001: "Something went wrong",
	2002: "Email has already been used",
}

func ErrorMessage(code int) string {
	return errorCodeToMessage[code]
}

func CustomError(code int) error {
	return fmt.Errorf("%d", code)
}
