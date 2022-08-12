package mysql

import (
	"RedBubble/models"
	"errors"
	"gorm.io/gorm"
)

var (
	ErrorCategoryEqualsNil = errors.New("没有分类")
	ErrorInvalidId         = errors.New("无效的分类id")
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

//获取某个分类详情
func GetCategoryById(id int64) (categoryDetail *models.Category, err error) {
	// SELECT gorm_id, category_name, introduction, gorm_created_at FROM category WHERE id = id ORDER BY id LIMIT 1;
	result := mdb.Select("gorm_id", "category_name", "introduction", "gorm_created_at").Where("gorm_id = ?", id).First(&categoryDetail)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		//没有分类
		err = ErrorInvalidId
	} else {
		//获取成功
		err = nil
	}
	return categoryDetail, err
}
