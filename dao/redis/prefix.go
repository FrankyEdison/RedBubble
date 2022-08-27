package redis

const (
	Prefix                = "redbubble:" // 项目key前缀
	KeyPostTimeZSet       = "post:time"  // zset类型，score记录发帖时间，member记录帖子
	KeyPostScoreZSet      = "post:score" // zset类型，score记录（点赞数-点灭数），member记录贴子
	KeyPostVoteZSetPrefix = "post:vote:" // zset类型，score记录用户id，member记录用户点赞/点灭情况; 前缀是post:vote:post_id

	KeyCommunitySetPrefix = "category:" // set;保存每个分区下帖子的id
)

// 给redis key加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}
