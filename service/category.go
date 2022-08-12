package service

import (
	"RedBubble/dao/mysql"
	"RedBubble/models"
)

//获取所有帖子分类
func GetAllCategory() (categories []*models.Category, err error) {
	return mysql.GetAllCategory()
}

//获取某个分类详情
func GetCategoryById(id int64) (categoryDetail *models.Category, err error) {
	return mysql.GetCategoryById(id)
}
