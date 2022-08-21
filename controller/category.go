package controller

import (
	"RedBubble/common/response"
	"RedBubble/common/responseCode"
	"RedBubble/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

//获取所有帖子分类
func GetAllCategoryHandler(c *gin.Context) {
	//业务处理，切片数据结构
	categories, err := service.GetAllCategory()
	if err != nil {
		zap.L().Error("获取所有帖子分类失败", zap.Error(err))
		response.Error(c, responseCode.CodeServerBusy)
		return
	}
	//响应
	response.Success(c, categories)
}

//获取某个分类详情
func GetCategoryByIdHandler(c *gin.Context) {
	// 1. 获取参数（在请求路径里的分类id）
	idStr := c.Param("cateId")
	id, err := strconv.ParseInt(idStr, 10, 64) //10进制，64位
	if err != nil {
		response.Error(c, responseCode.CodeInvalidParam)
		return
	}

	// 2. 根据id获取分类详情
	categoryDetail, err := service.GetCategoryById(id)
	if err != nil {
		zap.L().Error("获取分类详情失败", zap.Error(err))
		response.Error(c, responseCode.CodeServerBusy)
		return
	}
	response.Success(c, categoryDetail)
}
