package redis

import (
	"errors"
	"fmt"
	"forum-server/global"
	"forum-server/model/forum"
	"forum-server/utils"
	"math"
	"strconv"
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
	//pipeline1.Expire(ctx, getRedisKey(postId), time.Hour*24*7)
	_, err := pipeline1.Exec(ctx)
	if err != nil {
		return err
	}

	// 现在点赞value = -1 之前点赞ov = 1 看是从-1到1还是从1到-1op = -1 diff = 2 ex: value = 1 ov = -1 op = 1 diff = 2
	var op float64
	var fsd forum.FrmStarDetailString
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
			"post_id", postId,
			"user_id", userId,
			"status", value,
			"created_at", time.Now().Format("2006-01-02 15:04:05"),
			"updated_at", 0)
	} else {
		pipeline.HSet(ctx, getVoteRedisKey(postId, userId), fsd.UpdatedAt, time.Now().Format("2006-01-02 15:04:05"))
	}
	// 维护一个post_counter
	var dif int64
	if value == 1 {
		dif = 1
	} else if (value == 0 && ov != -1) || (value == -1 && ov == 1) {
		dif = -1
	} else {
		dif = 0
	}
	if pipeline.HExists(ctx, getRedisKey(KeyPostLikedCounterHSetPF), postId).Val() {
		pipeline.HSet(ctx, getRedisKey(KeyPostLikedCounterHSetPF), postId, dif)
	} else {
		pipeline.HIncrBy(ctx, getRedisKey(KeyPostLikedCounterHSetPF), postId, dif)
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
	var f forum.FrmStarDetailString
	var likeNum int64
	// 循环从post_set中pop出postId直到空
	for global.GVA_REDIS.SCard(ctx, getRedisKey(KeyPostLikedSetPF)).Val() != 0 {
		postId := global.GVA_REDIS.SPop(ctx, getRedisKey(KeyPostLikedSetPF)).Val()
		// 根据postId 每次从like_set中pop出一个userId直到空
		//pipeline := global.GVA_REDIS.Pipeline()
		for global.GVA_REDIS.SCard(ctx, getRedisKey(postId)).Val() != 0 {
			userId := global.GVA_REDIS.SPop(ctx, getRedisKey(postId)).Val()
			err = global.GVA_REDIS.HMGet(ctx, getVoteRedisKey(postId, userId), "user_id", "post_id", "status", "create_at", "update_at").Scan(&f)
			if err != nil {
				fmt.Println(err)
				return err
			}
			status, _ := strconv.Atoi(f.Status)
			fsd := &forum.FrmStarDetail{
				StarId:      utils.GenID(),
				LikedPostId: f.PostId,
				LikedUserId: f.UserId,
				Status:      int8(status),
				CreatedAt:   utils.TimeStringToGoTime(f.CreatedAt, utils.TimeTemplates),
				UpdatedAt:   utils.TimeStringToGoTime(f.UpdatedAt, utils.TimeTemplates),
			}
			// 将点赞信息更新进mysql
			var out forum.FrmUserStar
			resultFind := db.Table("frm_user_stars").
				Select("star_id").
				Where("liked_post_id = ? and liked_user_id = ?", postId, userId).Find(&out)
			if resultFind.RowsAffected < 1 {
				db.Table("frm_user_stars").Create(fsd)
			} else {
				db.Table("frm_user_stars").
					Where("liked_post_id = ? and liked_user_id = ?", postId, userId).
					Updates(forum.FrmStarDetail{Status: int8(status),
						UpdatedAt: utils.TimeStringToGoTime(f.UpdatedAt, utils.TimeTemplates)})
			}
			//pipeline.Del(ctx, getVoteRedisKey(postId, userId))
			// 将点赞数量更新进mysql
			db.Table("frm_posts").
				Select("like_num").
				Where("post_id = ?", postId).
				Find(&likeNum)
			newLikeNum, _ := strconv.ParseInt(global.GVA_REDIS.HGet(ctx, getRedisKey(KeyPostLikedCounterHSetPF), postId).Val(), 10, 64)
			if postId != "" {
				db.Table("frm_posts").
					Where("post_id = ?", postId).
					Update("like_num", likeNum+newLikeNum)
			}
			global.GVA_REDIS.HSet(ctx, getRedisKey(KeyPostLikedCounterHSetPF), postId, 0)
		}
	}
	return err
}
