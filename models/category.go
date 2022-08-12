package models

import "gorm.io/gorm"

type Category struct {
	Model        gorm.Model `gorm:"embedded;embeddedPrefix:gorm_"`                         // 说明是嵌套字段，并添加嵌套前缀
	CategoryName string     `gorm:"type:varchar(200);unique;index;not null;comment:帖子分类名"` // 指定在数据库中的字段类型，唯一，添加索引，非空
	Introduction string     `gorm:"type:varchar(1000);default:null;comment:分类介绍"`          // 字段类型，默认为空
}
