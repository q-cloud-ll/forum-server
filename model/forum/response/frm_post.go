package response

import (
	"forum/model/forum"
)

type FrmPostDetail struct {
	AuthorName          string `json:"author_name"`
	Avatar              string `json:"avatar"`
	VoteNum             int64  `json:"vote_num"`
	*forum.FrmPost      `json:"post"`
	*forum.FrmCommunity `json:"community"`
}
