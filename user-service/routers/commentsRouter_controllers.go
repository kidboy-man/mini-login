package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["user-service/controllers:UserPublicController"] = append(beego.GlobalControllerRouter["user-service/controllers:UserPublicController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("params", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user-service/controllers:UserPublicController"] = append(beego.GlobalControllerRouter["user-service/controllers:UserPublicController"],
        beego.ControllerComments{
            Method: "Register",
            Router: "/register",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("params", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

}
