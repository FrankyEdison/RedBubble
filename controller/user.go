package controller

import (
	"RedBubble/dao/mysql"
	"RedBubble/models"
	"RedBubble/service"
	"RedBubble/utils/validator_"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 1、用户注册
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数不是json格式，响应错误
		zap.L().Error("用户注册含非法参数", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//若不是validator的错误类型，随便返回就行
			ResponseError(c, CodeInvalidParam)
			return
		}
		//若是validator的错误类型，翻译一下错误再响应给前端
		ResponseErrorWithMsg(c, CodeInvalidParam, validator_.RemoveTopStruct(errs.Translate(validator_.Trans)))
		return
	}
	// 2. 业务处理
	if err := service.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}
