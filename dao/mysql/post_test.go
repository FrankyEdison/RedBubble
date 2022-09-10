package mysql

import (
	"RedBubble/models"
	"RedBubble/setting"
	"fmt"
	"testing"
)

// 实际上就是先初始化数据库连接，然后写死参数、调用dao层的方法

// 需要init的原因：mdb定义的时候就只是一个空值的*gorm.DB对象，不init的话它就一直是空值的*gorm.DB对象
func init() {
	dbCfg := setting.MySQLConfig{
		Host:         "127.0.0.1",
		Port:         3306,
		Username:     "root",
		Password:     "root",
		DBName:       "red_bubble",
		MaxIdleConns: 10,
		MaxOpenConns: 20,
	}
	err := Init(&dbCfg)
	if err != nil {
		fmt.Printf("初始化mysql连接失败, err:%v\n", err)
		return
	}
}

func TestAddPost(t *testing.T) {
	post := models.Post{
		PostId:     12466423381102593,
		UserId:     2325840394194944,
		Username:   "123",
		CategoryId: 1,
		Status:     0,
		Title:      "叱咤乐坛生力军",
		Content:    "fatBoy跳舞好强",
	}
	err := AddPost(&post)
	if err != nil {
		t.Fatalf("CreatePost insert record into mysql failed, err:%v\n", err)
	}
	t.Logf("CreatePost insert record into mysql success")
}
