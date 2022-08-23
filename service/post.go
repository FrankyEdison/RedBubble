package service

import (
	"RedBubble/dao/mysql"
	"RedBubble/models"
	"RedBubble/utils/snowflake"
)

// 发布帖子
func AddPost(post *models.Post) (err error) {
	// 生成PostId
	post.PostId = snowflake.GenerateID()
	// 保存到数据库
	return mysql.AddPost(post)

	//err = redis.CreatePost(p.ID, p.CategoryId)
}

//获取帖子详情
func GetPostDetailById(postId int64) (postDetail *models.Post, err error) {
	return mysql.GetPostDetailById(postId)
}

//分页获取所有帖子
func GetPostListByPage(pageSize int, pageNumber int) (postListByPage []*models.Post, err error) {
	//todo:把每个帖子的分类名也查出来，或post表添加分类名字段
	return mysql.GetPostListByPage(pageSize, pageNumber)
}
