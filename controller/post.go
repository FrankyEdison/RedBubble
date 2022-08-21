package controller

import (
	"RedBubble/common/parseUser"
	"RedBubble/common/response"
	"RedBubble/common/responseCode"
	"RedBubble/models"
	"RedBubble/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//发布帖子
func AddPostHandle(c *gin.Context) {
	// 1. 处理请求参数
	p1 := new(models.ParamAddPost)
	if err := c.ShouldBindJSON(p1); err != nil {
		zap.L().Error("发布帖子的请求参数有误", zap.Error(err))
		response.Error(c, responseCode.CodeInvalidParam)
		return
	}
	//从gin.context中获取jwt中存放的userID、username
	userID, username, err := parseUser.GetCurrentUser(c)
	if err != nil {
		response.Error(c, responseCode.CodeNeedLogin)
		return
	}
	//构造post model实例
	post := new(models.Post)
	post.UserId = userID
	post.Username = username
	post.CommunityID = p1.CommunityID
	post.Status = 0
	post.Title = p1.Title
	post.Content = p1.Content

	// 2. 创建帖子
	if err := service.AddPost(post); err != nil {
		zap.L().Error("创建帖子失败", zap.Error(err))
		response.Error(c, responseCode.CodeServerBusy)
		return
	}

	// 3. 返回响应
	response.Success(c, nil)
}
