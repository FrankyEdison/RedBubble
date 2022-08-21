package models

import (
	"gorm.io/gorm"
)

//内存对齐，推荐将相同类型的字段放在一起
type Post struct {
	Model       gorm.Model `gorm:"embedded;embeddedPrefix:gorm_"`                  // 说明是嵌套字段，并添加嵌套前缀
	PostId      int64      `gorm:"type:bigint;unique;index;not null;comment:帖子id"` // 指定在数据库中的字段类型，唯一，添加索引，非空
	UserId      int64      `gorm:"type:bigint;index;not null;comment:用户id"`        // 指定在数据库中的字段类型，添加索引，非空
	Username    string     `gorm:"type:varchar(200);not null;comment:用户名"`         // 指定在数据库中的字段类型，非空
	Likes       int32      `gorm:"type:bigint;default:0;comment:点赞数量"`             // 字段类型，默认为0
	CommunityID int64      `gorm:"type:bigint;index;not null;comment:社区id"`        // 字段类型，添加索引，非空
	Status      int8       `gorm:"type:tinyint;not null;comment:帖子状态"`             // 字段类型，非空
	Title       string     `gorm:"type:varchar(255);not null;comment:帖子标题"`        // 字段类型，非空
	Content     string     `gorm:"type:varchar(1000);not null;comment:帖子内容"`       // 字段类型，非空
}
