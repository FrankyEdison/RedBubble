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
		/**
		已注册的用户：123:111222
		*/
	}

	categoryGroup := apiGroup.Group("/category") // 帖子分类路由组
	{
		//获取所有帖子分类
		categoryGroup.GET("/getAllCategory", controller.GetAllCategoryHandler)
		//获取某个分类详情
		categoryGroup.GET("/getCategoryById/:cateId", controller.GetCategoryByIdHandler)
	}

	postGroup := apiGroup.Group("/post") //帖子路由组
	{
		//发表帖子
		postGroup.POST("/addPost", middleware.JWTAuthMiddleware(), controller.AddPostHandler)
		//获取帖子详情
		postGroup.GET("/:postId", controller.GetPostDetailHandler)
		//分页获取所有帖子（根据发表时间排序）
		postGroup.GET("/getPostListByPageByTime", controller.GetPostListByPageByTimeHandler)
		//分页获取所有帖子（根据点赞得分排序）
		postGroup.GET("/getPostListByPageByScore", controller.GetPostListByPageByScoreHandler)
	}

	voteGroup := apiGroup.Group("/vote") //点赞/点灭路由组
	{
		// 点赞/灭或者取消点赞/灭帖子
		voteGroup.POST("/votePost", middleware.JWTAuthMiddleware(), controller.VotePostHandler)
	}

	//测试使用，须登录后才能请求该路由，已注册中间件JWTAuthMiddleware()
	r.GET("/test", middleware.JWTAuthMiddleware(), controller.TestAuthHandler) // localhost:8081/test

	return r
}
