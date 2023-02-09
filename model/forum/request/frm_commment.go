package request

import "forum-server/model/common/request"

type Comment struct {
	PostId  int64  `json:"post_id"`
	Content string `json:"content"`
	request.PageInfo
}
