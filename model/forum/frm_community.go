package forum

import (
	"forum-server/global"
	"time"
)

type (
	FrmCommunity struct {
		global.GVA_MODEL
		CommunityId   int64  `json:"community_id" gorm:"not null;unique;comment:社区id" json:"community_id,omitempty"`
		CommunityName string `json:"community_name" gorm:"not null;unique;comment:社区名" json:"community_name,omitempty"`
		Introduction  string `json:"introduction,omitempty" gorm:"not null;comment:社区介绍" json:"introduction,omitempty"`
	}
)

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
