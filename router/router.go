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

	apiGroup := r.Group("/api")          // 给所有路由添加前缀/api
	userGroup := apiGroup.Group("/user") // user路由组
	{
		//用户注册
		userGroup.POST("/signUp", controller.SignUpHandler) // localhost:8081/api/user/signUp
		//用户登录
		userGroup.POST("/signIn", controller.SignInHandler)
	}

	categoryGroup := apiGroup.Group("/category") // 帖子分类路由组
	{
		//获取所有帖子分类
		categoryGroup.GET("/getAllCategory", controller.GetAllCategoryHandler)
	}

	//测试使用，须登录后才能请求该路由，已注册中间件JWTAuthMiddleware()
	r.GET("/test", middleware.JWTAuthMiddleware(), controller.TestAuthHandler) // localhost:8081/test

	return r
}
