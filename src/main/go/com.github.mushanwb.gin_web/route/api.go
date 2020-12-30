package route

import (
	"gin_web/src/main/go/com.github.mushanwb.gin_web/controller"
	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(r *gin.Engine) {

	c := new(controller.BaseController)
	r.GET("/ping", c.Ping)

}
