package redis

import (
	"context"
	"github.com/go-redis/redis/v9"
	"strconv"
	"time"
)

/**
在redis中存储以下四种数据：
①redbubble:category:categoryId，属于redis的Set类型，保存该分类下的所有postId
②redbubble:post:vote:postId，属于redis的ZSet类型，ZSet的Score保存用户点赞/灭的操作，ZSet的Member保存userId
③redbubble:post:score，属于redis的ZSet类型，ZSet的Score保存该帖子的点赞得分，ZSet的Member保存postId
④redbubble:post:time，属于redis的ZSet类型，ZSet的Score保存该帖子的创建时间，ZSet的Member保存postId
*/

// 发布帖子，要将帖子数据保存到redis中
func CreatePost(postId, communityId int64) error {
	ctx := context.Background()
	//创建事务
	pipeline := rdb.TxPipeline()
	// 按帖子加到'按发布时间排序的zset'中
	pipeline.ZAdd(ctx, GetRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})

	// 按帖子加到'按点赞数排序的zset'中
	pipeline.ZAdd(ctx, GetRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()), //基准值就是发布时间，点一个赞续432秒，如果一天有200个赞，那可以续到第二天
		Member: postId,
	})

	// 把帖子id加到社区的set
	cKey := GetRedisKey(KeyCategorySetPrefix + strconv.Itoa(int(communityId)))
	pipeline.SAdd(ctx, cKey, postId)

	_, err := pipeline.Exec(ctx)
	return err
}

type postAndScore struct {
	PostId string
	Score  float64
}

// 分页查询查询该分段的postId
func GetPostIdsByPage(pageSize int, pageNumber int) (postDetail []*postAndScore, err error) {
	ctx := context.Background()
	key := GetRedisKey(KeyPostScoreZSet)
	startIndex := (pageNumber - 1) * pageSize
	endIndex := startIndex + pageSize - 1
	result, err := rdb.ZRevRangeWithScores(ctx, key, int64(startIndex), int64(endIndex)).Result() //降序，分页查redis的ZSet

	for _, z := range result {
		var p postAndScore
		p.PostId = z.Member.(string)
		p.Score = z.Score
		postDetail = append(postDetail, &p)
	}
	return postDetail, err
}

// 批量获取帖子赞成票数量（不计点灭票）
func GetPostsLike(postIds []string) (likes []int64, err error) {
	ctx := context.Background()
	likes = make([]int64, 0, len(postIds))

	//创建事务，将批量查询放进同一个事务中，查询时便是对redis的同一次请求，可减少RTT
	pipeline := rdb.Pipeline()
	for _, postId := range postIds {
		postKey := GetRedisKey(KeyPostVoteZSetPrefix + postId)
		pipeline.ZCount(ctx, postKey, "1", "1") //查询投了赞成票的票数
	}
	cmders, err := pipeline.Exec(ctx) //执行请求
	if err != nil {
		return nil, err
	}

	//转换类型并存到结果集中
	for _, cmder := range cmders {
		val := cmder.(*redis.IntCmd).Val()
		likes = append(likes, val)
	}
	return
}

//根据分类id分页获取所有帖子（根据点赞得分排序）
func GetPostListByCategoryId(cateId int64, pageSize int, pageNumber int) (postDetail []*postAndScore, err error) {
	ctx := context.Background()
	categoryScoreKey := GetRedisKey(KeyCategorySetPrefix + "score:" + strconv.Itoa(int(cateId)))
	categoryKey := GetRedisKey(KeyCategorySetPrefix + strconv.Itoa(int(cateId)))
	scoreKey := GetRedisKey(KeyPostScoreZSet)
	// 判断redbubble:category:score:1是否存在，这是一个ZSet
	if rdb.Exists(ctx, categoryScoreKey).Val() < 1 {
		// 若该key不存在，则根据 redbubble:category:1 表和 redbubble:post:score 表生成 redbubble:category:score:1 表，该表记录了分类1中的所有帖子，且以点赞-点灭数为score排序
		pipeline := rdb.Pipeline()
		keys := make([]string, 0, 2)
		keys = append(keys, categoryKey)
		keys = append(keys, scoreKey)
		rdb.ZInterStore(ctx, categoryScoreKey, &redis.ZStore{
			Keys:      keys,
			Aggregate: "MAX", // Can be SUM, MIN or MAX
		})
		pipeline.Expire(ctx, categoryScoreKey, 60*time.Second) //设置超时时间一分钟
		_, err := pipeline.Exec(ctx)
		if err != nil {
			return nil, err
		}
	}

	// 以下是分页查询redbubble:category:score:1
	startIndex := (pageNumber - 1) * pageSize
	endIndex := startIndex + pageSize - 1
	result, err := rdb.ZRevRangeWithScores(ctx, categoryScoreKey, int64(startIndex), int64(endIndex)).Result() //降序，分页查redis的ZSet

	for _, z := range result {
		var p postAndScore
		p.PostId = z.Member.(string)
		p.Score = z.Score
		postDetail = append(postDetail, &p)
	}
	return postDetail, err
}
