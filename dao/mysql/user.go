package mysql

import (
	"RedBubble/models"
	"errors"
	"gorm.io/gorm"
)

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
	//result := mdb.Create(user)
	result := mdb.Select("UserId", "Username", "Password").Create(user)
	return result.Error
}

// 3、根据用户名获取用户
func GetUserByUsername(username string) (user *models.User, err error) {
	result := mdb.Where("username = ?", username).First(&user) // SELECT * FROM user WHERE username = username ORDER BY id LIMIT 1;
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		//用户不存在
		err = ErrorUserNotExist
	} else {
		//该用户已存在
		err = nil
	}
	return user, err
}
