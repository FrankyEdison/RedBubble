package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v9"
	"math"
	"time"
)

// 基于用户投票的帖子排名算法：https://blog.csdn.net/sd19871122/article/details/78033018

// 本项目使用简化版的投票分数
// 投一票就加432分   86400/200  --> 200张赞成票可以给你的帖子续一天

/* 投票的几种情况：
direction=1时，有两种情况：
   	1. 之前没有投过票，现在投赞成票    --> 票数差值的绝对值：1  +432
   	2. 之前投反对票，现在改投赞成票    --> 票数差值的绝对值：2  +432*2
direction=0时，有两种情况：
   	1. 之前投过反对票，现在要取消投票  --> 票数差值的绝对值：1  +432
	2. 之前投过赞成票，现在要取消投票  --> 票数差值的绝对值：1  -432
direction=-1时，有两种情况：
   	1. 之前没有投过票，现在投反对票    --> 票数差值的绝对值：1  -432
   	2. 之前投赞成票，现在改投反对票    --> 票数差值的绝对值：2  -432*2

投票的限制：
每个贴子自发表之日起一个星期之内允许用户投票，超过一个星期就不允许再投票了。
	1. 到期之后将redis中保存的赞成票数及反对票数存储到mysql表中
	2. 到期之后删除 KeyPostVoteZSetPrefix
*/

/**
在redis中存储以下四种数据：
①redbubble:category:categoryId，属于redis的Set类型，保存该分类下的所有postId
②redbubble:post:vote:postId，属于redis的ZSet类型，ZSet的Score保存用户点赞/灭的操作，ZSet的Member保存userId
③redbubble:post:score，属于redis的ZSet类型，ZSet的Score保存该帖子的点赞得分，ZSet的Member保存postId
④redbubble:post:time，属于redis的ZSet类型，ZSet的Score保存该帖子的创建时间，ZSet的Member保存postId
*/

const (
	secondsOfOneWeekTime = 7 * 24 * 3600
	scorePerVote         = 432 // 每一票值多少分
)

var (
	ErrorVoteTimeExpire = errors.New("投票时间已过")
	ErrorVoteRepeated   = errors.New("不允许重复投票")
)

// 点赞/灭或者取消点赞/灭帖子
func VotePost(userId, postId string, voteAction float64) (err error) {
	ctx := context.Background()
	// 1. 判断投票限制
	// 去redis取帖子发布时间，点赞/灭时间只限帖子发布的一周内，若超过时间则不允许点赞/灭，且会将 点赞-点灭 的数量存到mysql中
	postTime := rdb.ZScore(ctx, GetRedisKey(KeyPostTimeZSet), postId).Val() // ZScore():获取元素的score
	if float64(time.Now().Unix())-postTime > secondsOfOneWeekTime {
		return ErrorVoteTimeExpire
	}

	// 2. 更新贴子的分数
	// 判断用户是否重复投票
	preVoteAction := rdb.ZScore(ctx, GetRedisKey(KeyPostVoteZSetPrefix+postId), userId).Val()
	if voteAction == preVoteAction {
		return ErrorVoteRepeated
	}

	var op float64
	if voteAction > preVoteAction {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(preVoteAction - voteAction) // 计算两次投票的差值

	//创建事务
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(ctx, GetRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postId) //增加元素分值

	// 3. 更新用户为该贴子点赞/灭的操作记录
	if voteAction == 0 {
		pipeline.ZRem(ctx, GetRedisKey(KeyPostVoteZSetPrefix+postId), userId) //删除元素
	} else {
		pipeline.ZAdd(ctx, GetRedisKey(KeyPostVoteZSetPrefix+postId), redis.Z{ //添加元素
			Score:  voteAction, // 赞成票还是反对票
			Member: userId,
		})
	}
	_, err = pipeline.Exec(ctx)
	return
}
