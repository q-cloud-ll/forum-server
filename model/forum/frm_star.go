package forum

import (
	"forum-server/global"
	"time"

	uuid "github.com/satori/go.uuid"
)

// FrmUserLike 用户点赞表
type FrmUserLike struct {
	global.GVA_MODEL
	LikedUserId uuid.UUID `json:"liked_user_id" gorm:"index;not null; comment:点赞的用户id"`
	LikedPostId int64     `json:"liked_post_id" gorm:"index;not null; comment:被点赞的帖子id"`
	Status      int8      `json:"status" gorm:"default:1; comment:点赞状态，1点赞，0取消，-1踩"`
}

type FrmStarDetail struct {
	PostId    int64     `json:"post_id"`
	UserId    uuid.UUID `json:"user_id"`
	Status    int8      `json:"status""`
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
}
