package forum

import (
	"forum-server/global"

	"github.com/google/uuid"
)

type FrmComment struct {
	global.GVA_MODEL
	CommentId int64     `json:"comment_id" gorm:"index;not null;unique;comment:评论id"`
	PostId    int64     `json:"post_id"`
	Type      int       `json:"type"   gorm:"size:5"`
	LikeNum   int       `json:"like_num"`
	ReplyId   int       `json:"reply_id"`
	UserId    uuid.UUID `json:"user_id"  gorm:"not null"`
	Content   string    `json:"content"   binding:"required"  gorm:"not null;type:longtext"`
}
