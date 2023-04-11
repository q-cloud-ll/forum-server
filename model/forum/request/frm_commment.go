package request

type PostComment struct {
	PostId  int64  `json:"post_id"`
	ReplyId int64  `json:"reply_id"`
	Pid     int64  `json:"pid"`
	Content string `json:"content"`
}

type CommentList struct {
	Page     int64  `json:"page" form:"page"`         // 页码
	PageSize int64  `json:"pageSize" form:"pageSize"` // 每页大小
	PostId   string `json:"post_id" form:"post_id"`
	Order    string `json:"order" form:"order" example:"score"`
}

type CdCommentList struct {
	Page      int64  `json:"page" form:"page"`         // 页码
	PageSize  int64  `json:"pageSize" form:"pageSize"` // 每页大小
	PostId    string `json:"post_id" form:"post_id"`
	CommentId string `json:"comment_id" form:"comment_id"`
	Order     string `json:"order" form:"order" example:"score"`
}

type StarComment struct {
	Pid       string `json:"pid"`
	PostId    string `json:"post_id"`
	CommentId string `json:"comment_id"`
	Direction int8   `json:"direction"`
}
