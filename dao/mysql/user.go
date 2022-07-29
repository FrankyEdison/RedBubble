package mysql

import (
	"RedBubble/models"
	"errors"
)

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)

//1、判断该用户名是否已存在
func CheckUsernameIsExist(username string) (err error) {

	return
}

//2、添加用户
func InsertUser(user *models.User) (err error) {

	return
}
