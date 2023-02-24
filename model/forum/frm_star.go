package forum

import (
	"forum/global"
	"time"
)

// FrmUserStar 用户点赞表
type FrmUserStar struct {
	global.GVA_MODEL
	StarId      int64  `json:"star_id" gorm:"index;not null;unique; comment:点赞id"`
	LikedUserId string `json:"liked_user_id" gorm:"index;not null; comment:点赞的用户id"`
	LikedPostId string `json:"liked_post_id" gorm:"index;not null; comment:被点赞的帖子id"`
	Status      int8   `json:"status" gorm:"default:1; comment:点赞状态，1点赞，0取消，-1踩"`
}

// FrmStarDetail 用户点赞细节
type FrmStarDetail struct {
	StarId      int64     `json:"star_id"`
	LikedUserId string    `json:"liked_user_id"`
	LikedPostId string    `json:"liked_post_id"`
	Status      int8      `json:"status"`
	CreatedAt   time.Time // 创建时间
	UpdatedAt   time.Time // 更新时间
}

type FrmStarDetailString struct {
	PostId    string `json:"post_id" redis:"post_id"`
	UserId    string `json:"user_id" redis:"user_id"`
	Status    string `json:"status" redis:"status"`
	CreatedAt string `json:"created_at" redis:"created_at"` // 创建时间
	UpdatedAt string `json:"updated_at" redis:"updated_at"` // 更新时间
}
