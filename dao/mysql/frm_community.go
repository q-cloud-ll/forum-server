package mysql

import (
	"forum/global"
	"forum/model/forum"
	"forum/model/forum/response"

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

// FrmGetCommunityInfo 根据参数获取社区信息
func FrmGetCommunityInfo(communityID int64) (data []response.FrmCommunityInfo, err error) {
	db := global.GVA_DB
	if communityID == 0 {
		err = db.Table("frm_communities").Find(&data).Error
	} else {
		err = db.Table("frm_communities").Where("community_id = ?", communityID).Find(&data).Error
	}
	return
}
