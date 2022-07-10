package repository

import (
	"auth-service/conf"
	models "auth-service/models/external"
	"auth-service/utils"
	"encoding/json"
	"fmt"
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

	payload, err := json.Marshal(user)
	if err != nil {
		return
	}

	_, err = utils.PostJSONRequest(url, string(payload))
	if err != nil {
		return
	}
	return
}
