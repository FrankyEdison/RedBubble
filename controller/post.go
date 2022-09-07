package controller

import (
	"RedBubble/common/parseUser"
	"RedBubble/common/response"
	"RedBubble/common/responseCode"
	"RedBubble/models"
	"RedBubble/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

//发布帖子
func AddPostHandler(c *gin.Context) {
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
	post.CategoryId = p1.CategoryId
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

//获取帖子详情
func GetPostDetailHandler(c *gin.Context) {
	// 1. 获取参数（在请求路径里的帖子id）
	idStr := c.Param("postId")
	id, err := strconv.ParseInt(idStr, 10, 64) //10进制，64位
	if err != nil {
		response.Error(c, responseCode.CodeInvalidParam)
		return
	}
	// 2.1 获取帖子详情
	postDetail, err := service.GetPostDetailById(id)
	if err != nil {
		zap.L().Error("获取帖子失败", zap.Error(err))
		response.Error(c, responseCode.CodeServerBusy)
		return
	}
	// 2.2 获取分类详情
	categoryDetail, err := service.GetCategoryById(postDetail.CategoryId)
	if err != nil {
		zap.L().Error("获取分类失败", zap.Error(err))
		response.Error(c, responseCode.CodeServerBusy)
		return
	}

	// 3.响应
	type ReturnDetail struct {
		//加了json字段的话就等于分了类，不然Post和Category里同名的字段就会糅合一起导致不显示
		*models.Post     `json:"post"`
		*models.Category `json:"category"`
	}
	p2 := &ReturnDetail{
		Post:     postDetail,
		Category: categoryDetail,
	}
	response.Success(c, p2)
}

//分页获取所有帖子（根据发表时间排序）
func GetPostListByPageByTimeHandler(c *gin.Context) {
	// 1. 处理请求参数
	p1 := new(models.ByPage)
	if err := c.ShouldBind(p1); err != nil {
		response.Error(c, responseCode.CodeInvalidParam)
		return
	}

	postListByPage, err := service.GetPostListByPageByTime(p1.PageSize, p1.PageNumber)
	if err != nil {
		zap.L().Error("分页获取帖子列表失败", zap.Error(err))
		response.Error(c, responseCode.CodeServerBusy)
		return
	}

	response.Success(c, postListByPage)
}

//分页获取所有帖子（根据点赞得分排序）
/**
因为点赞得分的数据是存在redis里的，数据类型是ZSet，ZSet的score存分数、member存postId。
步骤：
	1、先去redis查询该分段的postId
	2、根据postId去mysql查详细信息
问：为什么member不能存整个帖子的信息呢？这样直接查redis就可以查出来所有帖子的信息了
答：因为点赞得分会经常被修改，修改时是根据member来找到这个score从而修改score，member要完全一致才能找到对应的score，member太长的话会导致修改点赞得分不方便
*/
func GetPostListByPageByScoreHandler(c *gin.Context) {
	// 1. 处理请求参数
	p1 := new(models.ByPage)
	if err := c.ShouldBind(p1); err != nil {
		response.Error(c, responseCode.CodeInvalidParam)
		return
	}

	postListByPage, err := service.GetPostListByPageByScore(p1.PageSize, p1.PageNumber)
	if err != nil {
		zap.L().Error("分页获取帖子列表失败", zap.Error(err))
		response.Error(c, responseCode.CodeServerBusy)
		return
	}

	response.Success(c, postListByPage)
}
