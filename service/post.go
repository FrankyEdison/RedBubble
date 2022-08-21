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

	//err = redis.CreatePost(p.ID, p.CommunityID)
}

//获取帖子详情
func GetPostDetailById(postId int64) (postDetail *models.Post, err error) {
	return mysql.GetPostDetailById(postId)
}
