package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	math_rand "math/rand"
	database "user-service/databases"
	_ "user-service/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))

	database.InitDB()

	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	beego.Run()
}
