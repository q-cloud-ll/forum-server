package redis

import (
	"forum/global"
	frmReq "forum/model/forum/request"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

// FrmPostCreatePost 向redis存储帖子的key，做分数排行和时间排行
func FrmPostCreatePost(postId, communityId int64) error {
	pipeline := global.GVA_REDIS.TxPipeline()
	// 帖子时间
	pipeline.ZAdd(ctx, getRedisKey(KeyPostTimeZSet), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})

	pipeline.ZAdd(ctx, getRedisKey(KeyPostScoreZSet), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})

	// 将帖子id添加到社区的set 例如forum:community:3826223906033666 1944204179673088
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityId)))
	pipeline.SAdd(ctx, cKey, postId)
	_, err := pipeline.Exec(ctx)

	return err
}

// FrmGetPostIdsInOrder 从redis中获取id
func FrmGetPostIdsInOrder(p *frmReq.PostList) ([]string, error) {
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == frmReq.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}

	return getIdsFormKey(key, p.Page, p.PageSize)
}

// getIdsFormKey 根据时间或者得分，获取帖子id的排行
func getIdsFormKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1

	return global.GVA_REDIS.ZRevRange(ctx, key, start, end).Result()
}

// FrmGetCommunityPostIdsInOrder 根据社区查询帖子id
func FrmGetCommunityPostIdsInOrder(p *frmReq.PostList) ([]string, error) {
	// 看是否是时间排序还是分数排序
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == frmReq.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}
	// 社区的key
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityID)))
	// 利用缓存key减少 ZInterStore 执行次数
	key := orderKey + strconv.Itoa(int(p.CommunityID))
	// cKey中保存的是set集合社区对应的post_id,orderKey保存的是post_id对应的分值，ZInterStore保存的key是cKey和orderKey的交集，max是取他俩的分值最高，set默认分为0
	if global.GVA_REDIS.Exists(ctx, key).Val() < 1 {
		pipeline := global.GVA_REDIS.Pipeline()
		pipeline.ZInterStore(ctx, key, &redis.ZStore{
			Aggregate: "MAX",
			Keys:      []string{cKey, orderKey},
		})
		pipeline.Expire(ctx, key, 3*time.Second)
		_, err := pipeline.Exec(ctx)
		if err != nil {
			return nil, err
		}
	}
	return getIdsFormKey(key, p.Page, p.PageSize)
}
