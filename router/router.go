package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //设置成发布模式
	}
	r := gin.New()
	//注册中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	//注册
	r.POST("/signup", controller.SignUpHandler)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "404",
		})
	})
	return r
}
