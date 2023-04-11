package redis

import (
	"context"
	"errors"
)

var (
	ctx               = context.Background()
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated   = errors.New("不允许重复投票")
)

const (
	Prefix                    = "forum:"       // 项目key前缀
	KeyPostTimeZSet           = "post:time"    // zset;贴子及发帖时间
	KeyPostScoreZSet          = "post:score"   // zset;贴子及投票的分数
	KeyPostVotedZSetPF        = "post:voted:"  // zset;记录用户及投票类型;参数是post id
	KeyPostLikedSetPF         = "post:liked:"  // 放所有被点赞的帖子
	KeyPostLikedCounterHSetPF = "post:counter" // 储存每个帖子的counter

	KeyCommunitySetPF = "community:" // set;保存每个分区下帖子的id
	KeyVotedHSetPF    = "::"

	KeyCommentTimeZSetPF         = "comment:time:"  // zset;评论及发帖时间
	KeyCommentScoreZSetPF        = "comment:score:" // zset;评论及投票的分数
	KeyCommentVotedZSetPF        = "comment:voted:" // zset;记录用户及投票类型;参数是comment id
	KeyCommentChildrenNumSetPF   = "comment:children:"
	KeyCommentLikedCounterHSetPF = "comment:counter" // 储存每个帖子的counter

	KeyFollowersSetPF  = "followers:"
	KeyFolloweesZsetPF = "followees:" // zset

	KeyQRCodeTicket = "wxAccessToken"

	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 // 每一票值多少分
)

// 给redis key加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}

// getVoteRedisKey 获取投票的key
func getVoteRedisKey(postId string, userId string) string {
	return postId + KeyVotedHSetPF + userId
}
