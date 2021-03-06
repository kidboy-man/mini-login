package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["user-service/controllers:UserAdminController"] = append(beego.GlobalControllerRouter["user-service/controllers:UserAdminController"],
        beego.ControllerComments{
            Method: "GetUsers",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("limit"),
				param.New("page"),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user-service/controllers:UserAdminController"] = append(beego.GlobalControllerRouter["user-service/controllers:UserAdminController"],
        beego.ControllerComments{
            Method: "GetUser",
            Router: "/:userID",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("userID", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user-service/controllers:UserAdminController"] = append(beego.GlobalControllerRouter["user-service/controllers:UserAdminController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: "/:userID",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(
				param.New("userID", param.IsRequired, param.InPath),
				param.New("params", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user-service/controllers:UserAdminController"] = append(beego.GlobalControllerRouter["user-service/controllers:UserAdminController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: "/:userID",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(
				param.New("userID", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user-service/controllers:UserInternalController"] = append(beego.GlobalControllerRouter["user-service/controllers:UserInternalController"],
        beego.ControllerComments{
            Method: "CreateUser",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("params", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user-service/controllers:UserPrivateController"] = append(beego.GlobalControllerRouter["user-service/controllers:UserPrivateController"],
        beego.ControllerComments{
            Method: "GetMyProfile",
            Router: "/my",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["user-service/controllers:UserPrivateController"] = append(beego.GlobalControllerRouter["user-service/controllers:UserPrivateController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: "/my",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(
				param.New("params", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

}
