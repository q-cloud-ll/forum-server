package redis

import (
	"forum/global"
	"time"

	"github.com/go-redis/redis/v8"
)

// AddFollower 添加关注者集合   followerId 关注者 followeeId被关注者
// er是走我们去吹风 ee 是cherry 我去点cherry的关注 添加我的的关注列表
func AddFollower(followeeId, followerId string) (err error) {
	return global.GVA_REDIS.SAdd(ctx, getRedisKey(KeyFollowersSetPF+followerId), followeeId).Err()
}

// UnFollower 从关注列表移除
func UnFollower(followeeId, followerId string) (err error) {
	return global.GVA_REDIS.SRem(ctx, getRedisKey(KeyFollowersSetPF+followerId), followeeId).Err()
}

// AddFollowee 添加粉丝列表有序集合
func AddFollowee(followerId, followeeId string) (err error) {
	return global.GVA_REDIS.ZAdd(ctx, getRedisKey(KeyFolloweesZsetPF+followeeId), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: followerId,
	}).Err()
}

// UnFollowee 移除粉丝
func UnFollowee(followerId, followeeId string) (err error) {
	return global.GVA_REDIS.ZRem(ctx, getRedisKey(KeyFolloweesZsetPF+followeeId), followerId).Err()
}

// GetFollowees 获取一个用户的粉丝列表
func GetFollowees(userId string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1
	followees, err := global.GVA_REDIS.ZRevRange(ctx, getRedisKey(KeyFolloweesZsetPF+userId), start, end).Result()
	if err != nil {
		return nil, err
	}
	var followerIds []string
	for _, followee := range followees {
		followerIds = append(followerIds, followee)
	}

	return followerIds, err
}

func GetFollowers(userId string) ([]string, error) {
	return global.GVA_REDIS.SMembers(ctx, getRedisKey(KeyFollowersSetPF+userId)).Result()
}
