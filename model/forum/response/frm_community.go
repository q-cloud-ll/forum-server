package response

import (
	"time"

	"gorm.io/gorm"
)

// FrmCommunityInfo 社区信息
type FrmCommunityInfo struct {
	CommunityId   string         `json:"community_id"`
	CommunityName string         `json:"community_name"`
	Introduction  string         `json:"introduction,omitempty"`
	CreatedAt     time.Time      `json:"created_at"` // 创建时间
	UpdatedAt     time.Time      `json:"updated_at"` // 更新时间
	DeletedAt     gorm.DeletedAt `json:"-"`          // 删除时间
}
