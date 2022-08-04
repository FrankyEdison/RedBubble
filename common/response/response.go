package response

import (
	"RedBubble/common/responseCode"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
{
	"code": 1000, // 程序中的错误码
	"msg": xx,     // 提示信息
	"data": {},    // 数据
}
*/
//通用响应对象

type ResponseData struct {
	Code responseCode.ResCode `json:"code"`
	Msg  interface{}          `json:"msg"`
	Data interface{}          `json:"data"`
}

func Error(c *gin.Context, code responseCode.ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ErrorWithMsg(c *gin.Context, code responseCode.ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: responseCode.CodeSuccess,
		Msg:  responseCode.CodeSuccess.Msg(),
		Data: data,
	})
}
