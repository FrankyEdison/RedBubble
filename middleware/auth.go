package middleware

import (
	"RedBubble/common/parseUser"
	"RedBubble/common/response"
	"RedBubble/common/responseCode"
	"RedBubble/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/big"
	"strconv"
	"strings"
)

/**
本中间件的作用的：从请求头的Authorization里获取jwt，若没有携带jwt请求或者jwt不合法，则响应“需要登录”；若jwt合法则解析出当前登录用户的ID
               并将userId保存到c *gin.Context的上下文中，后续的处理函数能获得该值
*/

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 推荐Token放在Header的Authorization中，并使用Bearer开头
		// Authorization: Bearer xxxxxxx.xxx.xxx  / X-TOKEN: xxx.xxx.xx
		authHeader := c.Request.Header.Get("Authorization")
		//判断是否已携带Authorization
		fmt.Printf("authHeaer:%s\n", authHeader)
		if authHeader == "" {
			response.Error(c, responseCode.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格分割，将Authorization分成Bearer和jwt字符串两部分
		parts := strings.SplitN(authHeader, " ", 2)
		//判断Authorization格式是否准确
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Error(c, responseCode.CodeInvalidToken)
			c.Abort()
			return
		}
		// 获取parts[1]即jwt字符串，并解析它，须转换数据类型
		claims, err := jwt.ParseToken(parts[1])

		//解析出来的claims["user_id"]是float64的科学计数法类型，我们需要转换成int64类型
		userIdTest := claims["user_id"].(float64)
		//先将float64的科学计数法类型转成string类型
		newNum := big.NewRat(1, 1)
		newNum.SetFloat64(userIdTest)
		//再将string类型转成int64类型
		parseInt, _ := strconv.ParseInt(newNum.FloatString(0), 10, 64) //10进制，64位精度
		claims["user_id"] = parseInt

		//解析出来的claims["exp"]是float64的科学计数法类型，我们需要转换成int64类型
		expTest := claims["exp"].(float64)
		//先将float64的科学计数法类型转成string类型
		newNum.SetFloat64(expTest)
		//再将string类型转成int64类型
		parseInt, _ = strconv.ParseInt(newNum.FloatString(0), 10, 64) //10进制，64位精度
		claims["exp"] = parseInt

		fmt.Printf("claims:%v\n", claims)
		if err != nil {
			response.Error(c, responseCode.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set(parseUser.CtxUserIDKey, claims["user_id"])

		c.Next() // 后续的处理请求的函数中 可以用过c.Get(CtxUserIDKey) 来获取当前请求的用户信息
	}
}
