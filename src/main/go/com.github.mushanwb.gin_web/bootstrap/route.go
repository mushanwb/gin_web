package bootstrap

import (
	"gin_web/src/main/go/com.github.mushanwb.gin_web/route"
	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	r := gin.Default()
	route.RegisterApiRoutes(r)
	return r
}
