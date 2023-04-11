package redis

import (
	"forum/global"
	frmReq "forum/model/forum/request"
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

// FrmCreateComment 向redis存储帖子的key，做分数排行和时间排行
func FrmCreateComment(communityId, postId, pid int64) error {
	var pidStr, postIdStr string
	postIdStr = strconv.FormatInt(postId, 10)
	if pid == 0 {
		pidStr = ""
	} else {
		pidStr = strconv.FormatInt(pid, 10)
	}
	pipeline := global.GVA_REDIS.TxPipeline()
	pipeline.ZAdd(ctx, getRedisKey(KeyCommentTimeZSetPF+postIdStr+pidStr), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: communityId,
	})
	pipeline.ZAdd(ctx, getRedisKey(KeyCommentScoreZSetPF+postIdStr+pidStr), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: communityId,
	})
	_, err := pipeline.Exec(ctx)

	return err
}

// FrmGetCommentsIdsInOrder 从redis中获取id
func FrmGetCommentsIdsInOrder(p *frmReq.CommentList) ([]string, error) {
	key := getRedisKey(KeyCommentTimeZSetPF + p.PostId)
	if p.Order == frmReq.OrderScore {
		key = getRedisKey(KeyCommentScoreZSetPF + p.PostId)
	}

	return getIdsFormKey(key, p.Page, p.PageSize)
}

// FrmGetCdCommentsIdsInOrder 从redis中获取id
func FrmGetCdCommentsIdsInOrder(p *frmReq.CdCommentList) ([]string, error) {
	key := getRedisKey(KeyCommentTimeZSetPF + p.PostId + p.CommentId)
	if p.Order == frmReq.OrderScore {
		key = getRedisKey(KeyCommentScoreZSetPF + p.PostId + p.CommentId)
	}

	return getIdsFormKey(key, p.Page, p.PageSize)
}

// FrmIncrChildrenNum 如果写入的是子评论，则父评论的子评论数量++
func FrmIncrChildrenNum(commentId int64) (err error) {
	if global.GVA_REDIS.Exists(ctx, getRedisKey(KeyCommentChildrenNumSetPF+strconv.FormatInt(commentId, 10))).Val() < 1 {
		err = global.GVA_REDIS.Set(ctx, getRedisKey(KeyCommentChildrenNumSetPF+strconv.FormatInt(commentId, 10)), 1, 0).Err()
		return err
	}
	err = global.GVA_REDIS.Incr(ctx, getRedisKey(KeyCommentChildrenNumSetPF+strconv.FormatInt(commentId, 10))).Err()
	return
}

// FrmGetChildrenNum 获取子评论数量
func FrmGetChildrenNum(commentId int64) (res int64, err error) {
	num, err := global.GVA_REDIS.Get(ctx, getRedisKey(KeyCommentChildrenNumSetPF+strconv.FormatInt(commentId, 10))).Result()
	if err != nil {
		return 0, err
	}
	res, _ = strconv.ParseInt(num, 10, 64)
	return res, err
}

// FrmStarComment 点赞评论
func FrmStarComment(uid, commentId, postId, pid string, dir float64) (err error) {
	// 更新帖子分数，先查当前用户给当前帖子的投票记录
	ov := global.GVA_REDIS.ZScore(ctx, getRedisKey(KeyCommentVotedZSetPF+commentId), commentId).Val()

	// 如果这次投票的值和之前保存的值一致，就提示不允许重复投票
	if dir == ov {
		return ErrVoteRepeated
	}
	// 现在点赞value = -1 之前点赞ov = 1 看是从-1到1还是从1到-1op = -1 diff = 2 ex: value = 1 ov = -1 op = 1 diff = 2
	var op float64
	if dir > op {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - dir)
	pipeline := global.GVA_REDIS.TxPipeline()
	// 记录用户点赞评论记录
	if pipeline.Exists(ctx, getVoteRedisKey(commentId, uid)).Val() < 1 {
		pipeline.HMSet(ctx, getVoteRedisKey(commentId, uid),
			"comment_id", commentId,
			"user_id", uid,
			"status", dir,
			"created_at", time.Now().Format("2006-01-02 15:04:05"),
			"updated_at", 0)
	} else {
		pipeline.Del(ctx, getVoteRedisKey(commentId, uid))
	}
	// 维护一个post_counter
	var dif int64
	if dir == 1 {
		dif = 1
	} else {
		dif = -1
	}
	// 记录每个评论的投票值
	if pipeline.HExists(ctx, getRedisKey(KeyCommentLikedCounterHSetPF), commentId).Val() {
		pipeline.HSet(ctx, getRedisKey(KeyCommentLikedCounterHSetPF), commentId, dif)
	} else {
		pipeline.HIncrBy(ctx, getRedisKey(KeyCommentLikedCounterHSetPF), commentId, dif)
	}
	// 计算zset排行分值
	pipeline.ZIncrBy(ctx, getRedisKey(KeyCommentScoreZSetPF+postId+pid), op*diff*scorePerVote, commentId)
	_, err = pipeline.Exec(ctx)
	return err
}

// FrmGetCommentStar 获取评论的点赞数量
func FrmGetCommentStar(commentId string) (likeNum int64, err error) {
	num, err := global.GVA_REDIS.HGet(ctx, getRedisKey(KeyCommentLikedCounterHSetPF), commentId).Result()
	if num == "" {
		return 0, nil
	}
	likeNum, _ = strconv.ParseInt(num, 10, 64)
	return likeNum, err
}
