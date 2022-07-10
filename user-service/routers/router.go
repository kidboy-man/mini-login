// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"user-service/controllers"
	"user-service/middlewares"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/user-service/v1",
		beego.NSNamespace("/admin/users",
			beego.NSBefore(middlewares.VerifyTokenAdmin),
			beego.NSInclude(
				&controllers.UserAdminController{},
			),
		),

		beego.NSNamespace("/private/users",
			beego.NSBefore(middlewares.VerifyToken),
			beego.NSInclude(
				&controllers.UserPrivateController{},
			),
		),

		beego.NSNamespace("/internal/users",
			beego.NSBefore(middlewares.IsPrivateIP),
			beego.NSInclude(
				&controllers.UserInternalController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
