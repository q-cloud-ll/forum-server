package redis

const (
	Prefix             = "forum:"      // 项目key前缀
	KeyPostTimeZSet    = "post:time"   // zset;贴子及发帖时间
	KeyPostScoreZSet   = "post:score"  // zset;贴子及投票的分数
	KeyPostVotedZSetPF = "post:voted:" // zset;记录用户及投票类型;参数是post id
	KeyPostVotedHSetPF = "::"
	KeyCommunitySetPF  = "community:" // set;保存每个分区下帖子的id
)

// 给redis key加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}

// getVoteRedisKey 获取投票的key
func getVoteRedisKey(postId string, userId string) string {
	return userId + KeyPostVotedHSetPF + postId
}
