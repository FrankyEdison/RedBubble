package controller

import (
	"RedBubble/common/response"
	"RedBubble/common/responseCode"
	"RedBubble/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// GetAllCategoryHandler 获取所有帖子分类
// @Summary 	获取所有帖子分类
// @Description 获取所有帖子分类
// @Tags 		帖子分类
// @Accept 		json
// @Produce 	json
// @Security 	ApiKeyAuth
// @Success 	200  {object}  models.ResponseAllCategory
// @Failure     500  {string}  string  "服务繁忙"
// @Router 		/category/getAllCategory [get]
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

// GetCategoryByIdHandler 获取某个分类详情
// @Summary 	获取某个分类详情
// @Description 根据分类id获取分类详情
// @Tags 		帖子分类
// @Accept 		json
// @Produce 	json
// @Param 		cateId  path  int  true  "分类id"
// @Success 	200  {object}  models.ResponseCategory
// @Failure     500  {string}  string  "请求参数错误"
// @Failure     500  {string}  string  "服务繁忙"
// @Router 		/category/getCategoryById/{cateId} [get]
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
