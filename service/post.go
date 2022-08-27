package service

import (
	"RedBubble/dao/mysql"
	"RedBubble/dao/redis"
	"RedBubble/models"
	"RedBubble/utils/snowflake"
	"go.uber.org/zap"
)

// 发布帖子
func AddPost(post *models.Post) (err error) {
	// 生成PostId
	post.PostId = snowflake.GenerateID()
	// 保存到mysql
	err = mysql.AddPost(post)
	if err != nil {
		zap.L().Error("保存到mysql失败", zap.Error(err))
		return err
	}
	// 保存到redis
	err = redis.CreatePost(post.PostId, post.CategoryId)
	if err != nil {
		zap.L().Error("保存到redis失败", zap.Error(err))
		return err
	}
	return
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
