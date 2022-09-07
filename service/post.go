package service

import (
	"RedBubble/dao/mysql"
	"RedBubble/dao/redis"
	"RedBubble/models"
	"RedBubble/utils/snowflake"
	"fmt"
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

//分页获取所有帖子（根据发表时间排序）
func GetPostListByPageByTime(pageSize int, pageNumber int) (postListByPage []*models.Post, err error) {
	//todo:把每个帖子的分类名也查出来，或post表添加分类名字段
	return mysql.GetPostListByPageByTime(pageSize, pageNumber)
}

//分页获取所有帖子（根据点赞得分排序）
func GetPostListByPageByScore(pageSize int, pageNumber int) (postListByPage []*models.Post, err error) {
	//1、先去redis分页查询查询该分段的postId和对应的分数
	postIdsWithScore, err := redis.GetPostIdsByPage(pageSize, pageNumber)
	if err != nil {
		zap.L().Error("分页查询redis失败", zap.Error(err))
		return nil, err
	}
	var postIds = make([]string, 0, len(postIdsWithScore)) //获取该分段的postId
	for _, value := range postIdsWithScore {
		postIds = append(postIds, value.PostId)
		fmt.Printf("postId:%s, score:%f\n", value.PostId, value.Score)
	}

	//2、根据postIds去mysql查详细信息
	//重点：返回的帖子列表须按照我传的postIds的顺序，不然就做不到得分从高到低
	postListByPage, err = mysql.GetPostListByPostIds(postIds)
	if err != nil {
		zap.L().Error("分页查询mysql失败", zap.Error(err))
		return nil, err
	}

	//3、给对应的post详情加上点赞得分
	for index, value := range postListByPage {
		value.Likes = int32((postIdsWithScore[index].Score - 1661593026) / 432)
		fmt.Printf("帖子标题:%s, score:%d\n", value.Title, value.Likes)
	}

	// PS:若只计点赞的票数，不计点灭的票数，则用以下方法
	// likes, err := redis.GetPostsLike(postIds)

	return postListByPage, nil
}
