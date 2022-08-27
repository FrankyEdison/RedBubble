package service

import (
	"RedBubble/dao/redis"
	"RedBubble/models"
	"go.uber.org/zap"
	"strconv"
)

// 点赞/灭或者取消点赞/灭帖子
func VotePost(userId int64, p *models.ParamVotePost) (err error) {
	zap.L().Debug("VoteForPost",
		zap.Int64("userId", userId),
		zap.String("PostId", p.PostId),
		zap.Int8("VoteAction", p.VoteAction))
	return redis.VotePost(strconv.Itoa(int(userId)), p.PostId, float64(p.VoteAction))
}
