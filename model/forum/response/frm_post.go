package response

import (
	"forum-server/model/forum"
)

type FrmPostDetail struct {
	AuthorName string `json:"author_name"`
	VoteNum    int64  `json:"vote_num"`
	*forum.FrmPost
	*forum.FrmCommunity `json:"community"`
}
