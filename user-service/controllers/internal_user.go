package controllers

import (
	"user-service/conf"
	"user-service/models"
	usecase "user-service/usecases"
)

// Operations about object
type UserInternalController struct {
	BaseController
	userUcase usecase.UserUsecase
}

func (c *UserInternalController) Prepare() {
	c.userUcase = usecase.NewUserUsecase(conf.AppConfig.DbClient)
}

// @Title Create User
// @Description Create User
// @Summary Create User
// @Success 200
// @Failure 403
// @Param params body models.User true "body of this request"
// @router / [post]
func (c *UserInternalController) CreateUser(params *models.User) *JSONResponse {
	err := c.userUcase.Create(params)
	return c.ReturnJSONResponse(nil, err)
}
