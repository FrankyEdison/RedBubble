package controller

import (
	"RedBubble/common/response"
	"RedBubble/common/responseCode"
	"RedBubble/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
