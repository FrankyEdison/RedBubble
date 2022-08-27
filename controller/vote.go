package controller

import (
	"RedBubble/common/parseUser"
	"RedBubble/common/response"
	"RedBubble/common/responseCode"
	"RedBubble/models"
	"RedBubble/service"
	"RedBubble/utils/validator_"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 点赞/灭或者取消点赞/灭帖子
func VotePostHandler(c *gin.Context) {
	//参数校验
	p := new(models.ParamVotePost)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) // 判断err是不是validator.ValidationErrors 类型
		if !ok {
			//若不是validator的错误类型，直接返回参数错误就行了
			response.Error(c, responseCode.CodeInvalidParam)
			return
		}
		//若是validator的错误类型，翻译一下错误再响应给前端
		response.ErrorWithMsg(c, responseCode.CodeInvalidParam, validator_.RemoveTopStruct(errs.Translate(validator_.Trans)))
		return
	}

	userId, _, err := parseUser.GetCurrentUser(c)
	if err != nil {
		response.Error(c, responseCode.CodeNeedLogin)
		return
	}

	if err := service.VotePost(userId, p); err != nil {
		zap.L().Error("点赞/点灭动作失败", zap.Error(err))
		response.Error(c, responseCode.CodeServerBusy)
		return
	}

	response.Success(c, nil)
}
