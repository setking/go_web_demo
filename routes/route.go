package routes

import (
	"github.com/gin-gonic/gin"
	"myApp/controller"
	"myApp/logger"
	"myApp/middlewares"
	"net/http"
)

func Setup(mode string) *gin.Engine {

	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	//注册
	r.POST("/signup", controller.SignUpHandler)
	//登录
	r.POST("/login", controller.LoginHandler)

	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		c.String(http.StatusOK, "ping")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
