package helpers

import (
	"auth-service/constants"
	"auth-service/datatransfers"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (hashedPassword string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		err = &datatransfers.CustomError{
			Code:    constants.HashPasswordInternalErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}

	hashedPassword = string(bytes)
	return
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
