package router

import (
	"RedBubble/controller"
	"RedBubble/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//用来测试路由能不能正常使用
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, viper.GetString("version")) //这里用viper才能实现热加载
	})

	//用户注册
	r.POST("/signUp", controller.SignUpHandler)

	//用户登录
	r.POST("/signIn", controller.SignInHandler)

	return r
}
