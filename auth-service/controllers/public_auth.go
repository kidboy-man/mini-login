package controllers

import (
	"auth-service/conf"
	"auth-service/datatransfers"
	usecase "auth-service/usecases"
)

// Operations about object
type AuthPublicController struct {
	BaseController
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
// @Param params body datatransfers.AuthRequest true "body of this request"
// @router /register [post]
func (c *AuthPublicController) Register(params *datatransfers.AuthRequest) *JSONResponse {
	err := c.authUcase.Register(params)
	return c.ReturnJSONResponse(nil, err)
}

// @Title Login
// @Description login
// @Summary login
// @Success 200
// @Failure 403
// @Param params body datatransfers.AuthRequest "body of this request"
// @router /login [post]
func (c *AuthPublicController) Login(params *datatransfers.AuthRequest) *JSONResponse {
	auth, err := c.authUcase.Login(params)
	return c.ReturnJSONResponse(auth, err)
}
