package repository

import (
	"auth-service/conf"
	"auth-service/constants"
	"auth-service/datatransfers"
	models "auth-service/models/external"
	"auth-service/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserRepository interface {
	Create(user *models.User) (err error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *models.User) (err error) {
	url := fmt.Sprintf(
		"%s/internal/users/",
		conf.AppConfig.UserServiceURL,
	)

	fmt.Println("url creating user:", url)

	payload, err := json.Marshal(user)
	if err != nil {
		err = &datatransfers.CustomError{
			Code:    constants.InternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}

	_, err = utils.PostJSONRequest(url, string(payload))
	if err != nil {
		return err
	}
	return nil
}
