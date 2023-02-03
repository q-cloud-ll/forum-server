package forum

import (
	"forum-server/dao/redis"
	frmReq "forum-server/model/forum/request"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

type VoteService struct{}

// FrmVotePost 帖子投票服务
func (voteService *VoteService) FrmVotePost(userId uuid.UUID, v *frmReq.FrmVoteData) (err error) {
	err = redis.FrmVotePost(userId.String(), strconv.FormatInt(v.PostId, 10), float64(v.Direction))
	return err
}
