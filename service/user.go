package service

import (
	"RedBubble/dao/mysql"
	"RedBubble/models"
	"RedBubble/utils/jwt"
	"RedBubble/utils/snowflake"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// md5的加密密码
const secret = "RedBubbleByFranky"

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
	fmt.Printf("userid=%d, username=%s, password=%s", user.UserId, user.Username, user.Password)
	// 4.使用md5对密码进行加密
	user.Password = encryptPassword(user.Password)
	// 5.保存进数据库
	return mysql.InsertUser(user)
}

// 1.4、使用go标准库提供的md5加密用户密码
func encryptPassword(originPassword string) (md5Password string) {
	//创建MD5算法
	h := md5.New()
	//写入待加密数据
	h.Write([]byte(secret))
	md5Password = hex.EncodeToString(h.Sum([]byte(originPassword)))
	return md5Password //16进制的字符串
}

// 2、用户登录
func SignIn(p *models.ParamSignIn) (token string, err error) {
	requestUser := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 1、根据用户名获取用户
	databaseUser, err := mysql.GetUserByUsername(requestUser.Username)
	if err != nil {
		// 用户不存在
		return "", err
	}
	// 2、判断用户输入的密码是否正确
	requestSecretPassword := encryptPassword(requestUser.Password)
	if requestSecretPassword != databaseUser.Password {
		// 用户名或密码错误
		return "", mysql.ErrorInvalidPassword
	}

	// 3、登录成功，返回JWT
	return jwt.GenerateToken(databaseUser.UserId, databaseUser.Username)
}
