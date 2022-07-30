package service

import (
	"RedBubble/dao/mysql"
	"RedBubble/models"
	"RedBubble/utils/snowflake"
)

// 1、用户注册
func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户存不存在，根据唯一用户名来判断
	if err = mysql.CheckUsernameIsExist(p.Username); err != nil {
		return err
	}
	// 2.生成用户ID
	userID := snowflake.GenerateID()
	// 3.构造一个User实例
	user := &models.User{
		UserId:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 4.保存进数据库
	return mysql.InsertUser(user)
}
