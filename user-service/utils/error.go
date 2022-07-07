package utils

import (
	"user-service/datatransfers"

	"gorm.io/gorm"
)

func IsErrRecordNotFound(err error) bool {
	v, ok := err.(*datatransfers.CustomError)
	if ok {
		return v.Error() == gorm.ErrRecordNotFound.Error()
	}

	return err == gorm.ErrRecordNotFound
}
