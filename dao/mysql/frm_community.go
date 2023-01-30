package mysql

import (
	"forum-server/global"
	"forum-server/model/forum"

	"gorm.io/gorm"
)

// FrmGetCommunityDetailById 根据社区id获取
func FrmGetCommunityDetailById(id int64) (community *forum.FrmCommunity, err error) {
	var cd *forum.FrmCommunity
	if err := global.GVA_DB.Table("frm_communities").
		Where("community_id = ? ", id).
		Find(&cd).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = ErrorInvalidID
		}
	}
	return cd, err
}
