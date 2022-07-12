package controllers

import (
	"user-service/conf"
	"user-service/datatransfers"
	"user-service/models"
	usecase "user-service/usecases"
)

// Operations about object
type UserPrivateController struct {
	BaseController
	userUcase usecase.UserUsecase
}

func (c *UserPrivateController) Prepare() {
	c.userUcase = usecase.NewUserUsecase(conf.AppConfig.DbClient)
}

// @Title Get My Profile
// @Description Get My Profile
// @Summary Get My Profile
// @Success 200
// @Failure 403
// @Param authorization header string true "bearer token in jwt"
// @router /my [get]
func (c *UserPrivateController) GetMyProfile() *JSONResponse {
	user, err := c.userUcase.GetByID(c.GetUserIDFromToken())
	return c.ReturnJSONResponse(user, err)
}

// @Title Update Profile
// @Description Update Profile
// @Summary Update Profile
// @Success 200
// @Failure 403
// @Param params body datatransfers.UpdateUserRequest true "body of this request"
// @Param authorization header string true "bearer token in jwt"
// @router /my [put]
func (c *UserPrivateController) UpdateUser(params *datatransfers.UpdateUserRequest) *JSONResponse {
	err := c.userUcase.Update(&models.User{
		ID:       c.GetUserIDFromToken(),
		FullName: params.FullName,
		Email:    params.Email,
	})

	return c.ReturnJSONResponse(nil, err)
}
