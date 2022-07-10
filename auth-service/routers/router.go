// @APIVersion 1.0.0
// @Title beego Auth-Service API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"auth-service/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/auth-service/v1",
		beego.NSNamespace("/public/auths",
			beego.NSInclude(
				&controllers.AuthPublicController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
