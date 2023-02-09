package forum

import (
	"forum-server/global"

	"github.com/google/uuid"
)

// FrmComment 评论表
type FrmComment struct {
	global.GVA_MODEL
	CommentId     int64     `json:"comment_id" gorm:"index;not null;unique;comment:评论id"`
	PostId        int64     `json:"post_id" gorm:"index;not null;comment:帖子id"`
	RootCommentId int64     `json:"root_comment_id" gorm:"index;not null;comment:根评论id"`
	ToCommentId   int64     `json:"to_comment_id" gorm:"index;comment:回复目标评论id"`
	LikeNum       int       `json:"like_num"`
	UserId        uuid.UUID `json:"user_id" gorm:"index;not null;comment:用户id"`
	Content       string    `json:"content" gorm:"not null"`
}
