package main

import (
	"gin_web/src/main/go/com.github.mushanwb.gin_web/bootstrap"
	_ "gin_web/src/main/go/com.github.mushanwb.gin_web/config"
)

func main() {
	r := bootstrap.SetupRoute()
	bootstrap.SetupDB()

	r.Run(":8080")
}
