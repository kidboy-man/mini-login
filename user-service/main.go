package main

import (
	database "user-service/databases"
	_ "user-service/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	database.InitDB()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
