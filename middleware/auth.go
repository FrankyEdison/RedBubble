package middleware

import (
	"RedBubble/controller"
	"RedBubble/utils/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 推荐Token放在Header的Authorization中，并使用Bearer开头
		// Authorization: Bearer xxxxxxx.xxx.xxx  / X-TOKEN: xxx.xxx.xx
		authHeader := c.Request.Header.Get("Authorization")
		//判断是否已携带Authorization
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格分割，将Authorization分成Bearer和jwt字符串两部分
		parts := strings.SplitN(authHeader, " ", 2)
		//判断Authorization格式是否准确
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		// 获取parts[1]即jwt字符串，并解析它
		claims, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set(controller.CtxUserIDKey, claims["user_id"])

		c.Next() // 后续的处理请求的函数中 可以用过c.Get(CtxUserIDKey) 来获取当前请求的用户信息
	}
}
