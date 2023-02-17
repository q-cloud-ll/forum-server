package response

import "time"

//type GetPostComment struct {
//	CommentSum []SignalPostComment
//}
//
//type SignalPostComment struct {
//	PostComment
//	ChildComments []PostComment
//}
//
//type ChildrenComment struct {
//	Comment []PostComment
//}

//type PostDetailCommentList struct {
//	Comments []*PostComment
//}

type PostComment struct {
	CommentId   int64     `json:"comment_id"`
	ChildrenNum int64     `json:"children_num"`
	LikeNum     int64     `json:"like_num"`
	Content     string    `json:"content"`
	CreateTime  time.Time `json:"createTime"`
	*PostCreate
	*PostReply
}

type PostCreate struct {
	CreateById   string `json:"create_by_id"`
	CreateByName string `json:"create_by_name"`
	CreateAvatar string `json:"create_avatar"`
}

type PostReply struct {
	ReplyUserId string `json:"reply_user_id"`
	ReplyAvatar string `json:"reply_avatar"`
	ReplyName   string `json:"replyName"`
}
