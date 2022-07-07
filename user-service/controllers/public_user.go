package controllers

import (
	"user-service/conf"
	"user-service/datatransfers"
	usecase "user-service/usecases"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about object
type UserPublicController struct {
	beego.Controller
	userUcase usecase.UserUsecase
}

func (c *UserPublicController) Prepare() {
	c.userUcase = usecase.NewUserUsecase(conf.AppConfig.DbClient)
}

// @Title Register
// @Description register
// @Summary register
// @Success 200
// @Failure 403
// @Param params body datatransfers.RegisterRequest true "body of this request"
// @router /register [post]
func (c *UserPublicController) Register(params *datatransfers.RegisterRequest) (response JSONResponse) {
	err := c.userUcase.Register(params)
	response.ReturnJSONResponse(nil, err)
	return
}

// @Title Login
// @Description login
// @Summary login
// @Success 200
// @Failure 403
// @Param params body models.User true "body of this request"
// @router /login [post]
func (c *UserPublicController) Login(params *datatransfers.LoginRequest) (response JSONResponse) {
	user, err := c.userUcase.Login(params)
	response.ReturnJSONResponse(user, err)
	return
}
