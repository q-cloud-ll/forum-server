package forum

import (
	"forum-server/dao/mysql"
	"forum-server/dao/redis"
	frmReq "forum-server/model/forum/request"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

type VoteService struct{}

// FrmVotePost 帖子投票服务
func (vs *VoteService) FrmVotePost(userId uuid.UUID, v *frmReq.FrmVoteData) (err error) {
	err = redis.FrmVotePost(userId.String(), strconv.FormatInt(v.PostId, 10), float64(v.Direction))
	return err
}

// GetPostVoteNum 获取帖子点赞数量
func GetPostVoteNum(postId string) (likeNum int64, err error) {
	cacheLikeNum, err := redis.FrmGetVoteNum(postId)
	dbNum, err := mysql.FrmPostVoteNum(postId)
	return cacheLikeNum + dbNum, err
}
