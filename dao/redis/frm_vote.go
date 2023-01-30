package redis

import (
	"errors"
	"forum-server/global"
	"forum-server/model/forum"
	"math"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 // 每一票值多少分
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated   = errors.New("不允许重复投票")
)

// FrmVotePost 给帖子投票
func FrmVotePost(userId, postId string, value float64) error {
	// 判断是否能投票，过了一周不能投票了
	postTime := global.GVA_REDIS.ZScore(ctx, getRedisKey(KeyPostTimeZSet), postId).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}

	// 更新帖子分数，先查当前用户给当前帖子的投票记录
	ov := global.GVA_REDIS.ZScore(ctx, getRedisKey(KeyPostVotedZSetPF+postId), userId).Val()

	// 如果这次投票的值和之前保存的值一致，就提示不允许重复投票
	if value == ov {
		return ErrVoteRepeated
	}
	// 现在点赞value = -1 之前点赞ov = 1 看是从-1到1还是从1到-1op = -1 diff = 2
	// value = 1 ov = -1 op = 1 diff = 2
	var op float64
	if value > op {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value)
	pipeline := global.GVA_REDIS.TxPipeline()
	pipeline.ZIncrBy(ctx, getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postId)
	if value == 1 {
		pipeline.SAdd(ctx, KeyPostVotedZSetPF+postId)
	}
	if global.GVA_REDIS.Exists(ctx, KeyPostVotedZSetPF+postId).Val() < 1 {
		fsd := &forum.FrmStarDetail{
			Status:    int8(value),
			CreatedAt: time.Now(),
		}
		pipeline.HSet(ctx, KeyPostVotedZSetPF+postId, KeyPostVotedZSetPF+userId, fsd)
	}

	if value == 0 {
		pipeline.ZRem(ctx, getRedisKey(KeyPostVotedZSetPF+postId), userId)
	} else {
		pipeline.ZAdd(ctx, getRedisKey(KeyPostVotedZSetPF+postId), &redis.Z{
			Score:  value,
			Member: userId,
		})
	}
	_, err := pipeline.Exec(ctx)
	return err
}

//func FrmStarPost(userId, postId string, value float64) error {
//
//}
