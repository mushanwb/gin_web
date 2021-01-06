package controller

import (
	"gin_web/src/main/go/com.github.mushanwb.gin_web/entity"
	"gin_web/src/main/go/com.github.mushanwb.gin_web/request"
	"github.com/gin-gonic/gin"
	"regexp"
)

func (*BaseController) Register(c *gin.Context) {
	var RegisterData request.RegisterRequestData

	if err := c.ShouldBind(&RegisterData); err != nil {
		entity.ReturnJson(c, 400, "参数校验错误", err)
		return
	}

	is, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, RegisterData.Phone)
	if !is {
		entity.ReturnJson(c, 400, "手机号格式错误", nil)
		return
	}

	entity.ReturnJson(c, 200, "注册成功", nil)
}
