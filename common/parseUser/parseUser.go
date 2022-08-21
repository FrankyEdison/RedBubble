package parseUser

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID"
const CtxUsernameKey = "username"

var ErrorUserNotLogin = errors.New("用户未登录")

// 从c *gin.Context的上下文中拿到auth中间件之前保存的userId和username
// 解析请求头jwt的部分已在auth中间件里完成，并将userId和username保存到c的上下文中
func GetCurrentUser(c *gin.Context) (userID int64, username string, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}

	uName, ok := c.Get(CtxUsernameKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	username, ok = uName.(string)
	if !ok {
		err = ErrorUserNotLogin
		return
	}

	return
}
