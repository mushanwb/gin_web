package entity

import "github.com/gin-gonic/gin"

func ReturnJson(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, gin.H{
		"message": message,
		"data":    data,
	})
	return
}
