package redis

import (
	"errors"
	"forum-server/global"
	"forum-server/model/forum"
	"math"
	"time"

	"gorm.io/gorm"

	"github.com/go-redis/redis/v8"
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

	// 创建两个集合, post_set, post_user_like_set
	pipeline1 := global.GVA_REDIS.Pipeline()
	// 无论value 是 0 1 -1都需要建立postId的set
	pipeline1.SAdd(ctx, getRedisKey(KeyPostLikedSetPF), postId)
	// new: 将被点赞的post当set的key，将所有给这个帖子点赞的user当作members
	pipeline1.SAdd(ctx, getRedisKey(postId), userId)
	_, err := pipeline1.Exec(ctx)
	if err != nil {
		return err
	}

	// 现在点赞value = -1 之前点赞ov = 1 看是从-1到1还是从1到-1op = -1 diff = 2 ex: value = 1 ov = -1 op = 1 diff = 2
	var op float64
	var fsd forum.FrmStarDetail
	if value > op {
		op = 1
	} else {
		op = -1
	}

	diff := math.Abs(ov - value)
	pipeline := global.GVA_REDIS.TxPipeline()
	// userId用户 -> postId帖子 -> value值（1，0，-1）,先判断用户是否对这个帖子投过票，现在这部分不是重复投票
	//pipeline.HSet(ctx, getRedisKey(KeyPostVotedZSetPF+postId), userId, fsd)
	if pipeline.Exists(ctx, getVoteRedisKey(postId, userId)).Val() < 1 {
		pipeline.HMSet(ctx, getVoteRedisKey(postId, userId),
			fsd.PostId, postId,
			fsd.UserId, userId,
			fsd.Status, value,
			fsd.CreatedAt, time.Now(),
			fsd.UpdatedAt, 0)
	} else {
		pipeline.HSet(ctx, getVoteRedisKey(postId, userId), fsd.UpdatedAt, time.Now())
	}
	// 维护一个post_counter
	if pipeline.Exists(ctx, getRedisKey(KeyPostLikedCounterHSetPF)).Val() < 1 {
		pipeline.HSet(ctx, getRedisKey(KeyPostLikedCounterHSetPF), postId, 0)
	} else {
		pipeline.HIncrBy(ctx, getRedisKey(KeyPostLikedCounterHSetPF), postId, int64(op*diff))
	}
	// 计算zset排行分值
	pipeline.ZIncrBy(ctx, getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postId)
	if value == 0 {
		pipeline.ZRem(ctx, getRedisKey(KeyPostVotedZSetPF+postId), userId)
	} else {
		pipeline.ZAdd(ctx, getRedisKey(KeyPostVotedZSetPF+postId), &redis.Z{
			Score:  value,
			Member: userId,
		})
	}
	_, err = pipeline.Exec(ctx)
	return err
}

// UpdateStarDetailFromRedisToMySQL 将点赞数据从缓存刷回数据库
func UpdateStarDetailFromRedisToMySQL(db *gorm.DB) (err error) {
	if db == nil {
		return errors.New("db Cannot be empty")
	}
	global.GVA_LOG.Info("定时任务执行")
	pipeline := global.GVA_REDIS.TxPipeline()
	for pipeline.Exists(ctx, getRedisKey(KeyPostLikedSetPF)).Val() < 1 {
		//postId := pipeline.SPop(ctx, getRedisKey(KeyPostLikedSetPF)).Val()

	}
	return err
}
