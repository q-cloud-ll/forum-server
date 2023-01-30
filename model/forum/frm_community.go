package forum

import (
	"forum-server/global"
	"time"
)

type FrmCommunity struct {
	global.GVA_MODEL
	CommunityId   int64  `json:"community" gorm:"not null;unique;comment:社区id"`
	CommunityName string `json:"community_name" gorm:"not null;unique;comment:社区名"`
	Introduction  string `json:"introduction,omitempty" gorm:"not null;comment:社区介绍"`
}

type CommunityDetail struct {
	CommunityId  int64  `json:"community_id"`
	Name         string `json:"name"`
	Introduction string `json:"introduction,omitempty"`
	CreateTime   time.Time
}

type Community struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
