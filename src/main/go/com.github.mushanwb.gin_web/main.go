package main

import (
	"gin_web/src/main/go/com.github.mushanwb.gin_web/bootstrap"
)

func main() {
	r := bootstrap.SetupRoute()
	r.Run(":8080")
}
