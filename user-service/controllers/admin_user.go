package controllers

import (
	"user-service/conf"
	"user-service/datatransfers"
	"user-service/models"
	usecase "user-service/usecases"
	"user-service/utils"
)

// Operations about object
type UserAdminController struct {
	BaseController
	userUcase usecase.UserUsecase
}

func (c *UserAdminController) Prepare() {
	c.userUcase = usecase.NewUserUsecase(conf.AppConfig.DbClient)
}

// @Title Get Users
// @Description Get Users
// @Summary Get Users
// @Success 200
// @Failure 403
// @Param limit query int false "limit of this request"
// @Param page query int false "page of this request"
// @router / [get]
func (c *UserAdminController) GetUsers(limit, page int) *JSONResponse {
	users, _, err := c.userUcase.GetAll(&datatransfers.ListQueryParams{
		Limit:  limit,
		Offset: utils.CalculateOffset(limit, page),
	})

	// TODO: set paginator
	return c.ReturnJSONResponse(users, err)
}

// @Title Get User Details
// @Description Get User Details
// @Summary Get User Details
// @Success 200
// @Failure 403
// @Param authorization header string true "bearer token in jwt"
// @Param userID path string true "user id"
// @router /:userID [get]
func (c *UserAdminController) GetUser(userID string) *JSONResponse {
	user, err := c.userUcase.GetByID(userID)
	return c.ReturnJSONResponse(user, err)
}

// @Title Update User
// @Description Update User
// @Summary Update User
// @Success 200
// @Failure 403
// @Param authorization header string true "bearer token in jwt"
// @Param userID path string true "user id"
// @Param params body datatransfers.UpdateUserRequest true "body of this request"
// @router /:userID [put]
func (c *UserAdminController) UpdateUser(userID string, params *datatransfers.UpdateUserRequest) *JSONResponse {
	err := c.userUcase.Update(&models.User{
		ID:       userID,
		FullName: params.FullName,
		Email:    params.Email,
	})

	return c.ReturnJSONResponse(nil, err)
}

// @Title Delete User
// @Description Delete User
// @Summary Delete User
// @Success 200
// @Failure 403
// @Param authorization header string true "bearer token in jwt"
// @Param userID path int true "user id"
// @router /:userID [delete]
func (c *UserAdminController) DeleteUser(userID string) *JSONResponse {
	err := c.userUcase.Delete(&models.User{
		ID: userID,
	})

	return c.ReturnJSONResponse(nil, err)
}
