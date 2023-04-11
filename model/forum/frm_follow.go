package forum

import "forum/global"

type FrmFollow struct {
	global.GVA_MODEL
	FollowId   int64  `json:"follow_id" gorm:"not null; index;comment:关注id"`
	FollowerId string `json:"follower_id" gorm:"not null; index"`
	FolloweeId string `json:"followee_id" gorm:"not null;index"`
}
