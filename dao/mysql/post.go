package mysql

import (
	"RedBubble/models"
)

// 新增帖子
func AddPost(post *models.Post) (err error) {
	result := mdb.Select("PostId", "UserId", "Username", "CategoryId", "Status", "Title", "Content").Create(post)
	return result.Error
}

// 获取帖子详情
func GetPostDetailById(postId int64) (postDetail *models.Post, err error) {
	result := mdb.Select("post_id", "username", "likes", "category_id", "status", "title", "content", "gorm_created_at").Where("post_id = ?", postId).First(&postDetail)
	return postDetail, result.Error
}

// 分页获取所有帖子
func GetPostListByPage(pageSize int, pageNumber int) (postListByPage []*models.Post, err error) {
	result := mdb.Select("post_id", "username", "likes", "category_id", "status", "title", "content", "gorm_created_at").Limit(pageSize).Offset((pageNumber - 1) * pageSize).Find(&postListByPage)
	return postListByPage, result.Error
}
