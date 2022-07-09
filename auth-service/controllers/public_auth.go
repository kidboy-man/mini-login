package controllers

import (
	"auth-service/conf"
	"auth-service/datatransfers"
	usecase "auth-service/usecases"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about object
type AuthPublicController struct {
	beego.Controller
	authUcase usecase.AuthUsecase
}

func (c *AuthPublicController) Prepare() {
	c.authUcase = usecase.NewAuthUsecase(conf.AppConfig.DbClient)
}

// @Title Register
// @Description register
// @Summary register
// @Success 200
// @Failure 403
// @Param params body datatransfers.RegisterRequest true "body of this request"
// @router /register [post]
func (c *AuthPublicController) Register(params *datatransfers.AuthRequest) (response JSONResponse) {
	err := c.authUcase.Register(params)
	response.ReturnJSONResponse(nil, err)
	return
}

// @Title Login
// @Description login
// @Summary login
// @Success 200
// @Failure 403
// @Param params body models.Auth true "body of this request"
// @router /login [post]
func (c *AuthPublicController) Login(params *datatransfers.AuthRequest) (response JSONResponse) {
	auth, err := c.authUcase.Login(params)
	response.ReturnJSONResponse(auth, err)
	return
}
