package mysql

import (
	"RedBubble/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// md5的加密密码
const secret = "RedBubbleByFranky"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)

//1、判断该用户名是否已存在（用户名是唯一的）
func CheckUsernameIsExist(username string) (err error) {
	var user models.User
	// Get first matched record. if no matched record, result.Error=gorm.ErrRecordNotFound. if matched record, result.Error=nil
	result := mdb.Where("username = ?", username).First(&user) // SELECT * FROM user WHERE username = username ORDER BY id LIMIT 1;
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		//用户名不存在
		err = nil
	} else {
		//该用户已存在
		err = ErrorUserExist
	}
	return err
}

//2、添加用户
func InsertUser(user *models.User) (err error) {
	fmt.Printf("userid=%d, username=%s, password=%s", user.UserId, user.Username, user.Password)
	// 使用md5对密码进行加密
	user.Password = encryptPassword(user.Password)
	//result := mdb.Create(user)
	result := mdb.Select("UserId", "Username", "Password").Create(user)
	return result.Error
}

//使用go标准库提供的md5加密用户密码
func encryptPassword(originPassword string) (md5Password string) {
	//创建MD5算法
	h := md5.New()
	//写入待加密数据
	h.Write([]byte(secret))
	md5Password = hex.EncodeToString(h.Sum([]byte(originPassword)))
	return md5Password //16进制的字符串
}
