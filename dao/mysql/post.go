package mysql

import (
	"RedBubble/models"
	"fmt"
	"strings"
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

// 分页获取所有帖子（根据发表时间排序）
func GetPostListByPageByTime(pageSize int, pageNumber int) (postListByPage []*models.Post, err error) {
	result := mdb.Select("post_id", "username", "likes", "category_id", "status", "title", "content", "gorm_created_at").Order("gorm_id desc").Limit(pageSize).Offset((pageNumber - 1) * pageSize).Find(&postListByPage)
	return postListByPage, result.Error
}

// 根据postIds批量获取帖子详情
func GetPostListByPostIds(postIds []string) (postListByPage []*models.Post, err error) {
	param := strings.Join(postIds, ",") //组装sql需要的参数
	fmt.Println("params:", param)
	//根据postIds的顺序查post详情，返回的结果的顺序跟postIds的顺序一样
	result := mdb.Raw("select post_id,username,likes,category_id,status,title,content,gorm_created_at from post where post_id in (" + param + ") order by FIELD(post_id," + param + ")").Find(&postListByPage)
	return postListByPage, result.Error
}
