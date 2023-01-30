package forum

import (
	"forum-server/global"

	uuid "github.com/satori/go.uuid"
)

type FrmPost struct {
	global.GVA_MODEL
	PostId      int64     `json:"post_id" gorm:"index;not null;unique;comment:帖子id"`
	CommunityId int64     `json:"community_id" gorm:"not null;comment:社区id"`
	CommentId   int64     `json:"comment_id"`
	ReplyId     int64     `json:"reply_id"`
	AuthorId    uuid.UUID `json:"author_id" gorm:"not null;comment:作者id"`
	Content     string    `json:"content" gorm:"type:longtext;comment:帖子内容"`
	Title       string    `json:"title" gorm:"size:500l;comment:帖子标题"`
	Type        int8      `json:"type" gorm:"size:5"`
	LikeNum     uint32    `json:"like_num" gorm:"default:0"`
	UnLikeNum   uint32    `json:"unLike_num" gorm:"default:0"`
}
