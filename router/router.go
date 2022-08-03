package router

import (
	"RedBubble/controller"
	"RedBubble/logger"
	"RedBubble/middleware"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//用户注册
	r.POST("/signUp", controller.SignUpHandler)

	//用户登录
	r.POST("/signIn", controller.SignInHandler)

	//测试使用，须登录后才能请求该路由，已注册中间件JWTAuthMiddleware()
	r.GET("/test", middleware.JWTAuthMiddleware(), controller.TestAuthHandler)

	return r
}
