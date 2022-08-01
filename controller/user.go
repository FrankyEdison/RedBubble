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
		zap.L().Error("用户注册含非法参数", zap.Error(err))   // 请求参数不是json格式，响应错误
		errs, ok := err.(validator.ValidationErrors) // 判断err是不是validator.ValidationErrors 类型
		if !ok {
			//若不是validator的错误类型，直接返回就行
			ResponseError(c, CodeInvalidParam)
			return
		}
		//若是validator的错误类型，翻译一下错误再响应给前端
		ResponseErrorWithMsg(c, CodeInvalidParam, validator_.RemoveTopStruct(errs.Translate(validator_.Trans)))
		return
	}
	// 2. 业务处理
	if err := service.SignUp(p); err != nil {
		zap.L().Error("注册失败,err:", zap.Error(err))
		// 用户已存在
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		// 服务器繁忙
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 注册成功，返回响应
	ResponseSuccess(c, nil)
}

// 2、用户登录
func SignInHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamSignIn)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("用户登录含非法参数", zap.Error(err))   // 请求参数不是json格式，响应错误
		errs, ok := err.(validator.ValidationErrors) // 判断err是不是validator.ValidationErrors 类型
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
	token, err := service.SignIn(p)
	if err != nil {
		zap.L().Error("登录失败，err：", zap.String("username", p.Username), zap.Error(err))
		// 用户不存在
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		// 用户名或密码错误
		if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		// 服务器繁忙
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 登录成功，返回响应
	ResponseSuccess(c, token)
}
