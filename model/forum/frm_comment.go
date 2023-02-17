package forum

import "forum-server/global"

// FrmComment 评论表
type FrmComment struct {
	global.GVA_MODEL
	CommentId int64  `json:"comment_id" gorm:"index;not null;unique;comment:评论id"`
	PostId    int64  `json:"post_id" gorm:"index;not null;comment:帖子id"`
	ReplyId   int64  `json:"reply_id" gorm:"index;not null"`
	Pid       int64  `json:"pid" gorm:"index;not null";comment:父id`
	LikeNum   int64  `json:"like_num"`
	UserId    string `json:"user_id" gorm:"index;not null;comment:用户id"`
	Content   string `json:"content" gorm:"not null"`
}

type FrmUserInfo struct {
	UserId   string `json:"user_id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}
