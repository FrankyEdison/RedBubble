package mysql

import (
	"RedBubble/models"
	"errors"
	"gorm.io/gorm"
)

var (
	ErrorCategoryEqualsNil = errors.New("没有分类")
)

//获取所有帖子分类
func GetAllCategory() (categories []*models.Category, err error) {
	// SELECT category_name, introduction FROM category;
	result := mdb.Select("category_name", "introduction").Find(&categories)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		//没有分类
		err = ErrorCategoryEqualsNil
	} else {
		//获取成功
		err = nil
	}
	return categories, err
}
