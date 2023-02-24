package mysql

import (
	"forum/global"
	"forum/model/forum"
	"forum/utils"

	"gorm.io/gorm/clause"
)

func AddFollower(followerId, followeeId string) (err error) {
	f := &forum.FrmFollow{
		FollowId:   utils.GenID(),
		FolloweeId: followeeId,
		FollowerId: followerId,
	}
	return global.GVA_DB.Model(&forum.FrmFollow{}).Create(f).Error
}

func UnFollowee(followerId, followeeId string) (err error) {
	return global.GVA_DB.Where("follower_id = ? and followee_id = ?", followerId, followeeId).Delete(&forum.FrmFollow{}).Error
}

// GetFollowUserInfo 获取关注列表
func GetFollowUserInfo(ids []string, page, pageSize int) (userInfo []*forum.UserInfo, err error) {
	limit := pageSize
	offset := pageSize * (page - 1)
	err = global.GVA_DB.Table("frm_users").
		Where("user_id in (?)", ids).
		Clauses(clause.OrderBy{
			Expression: clause.Expr{
				SQL:                "FIELD(user_id, ?)",
				Vars:               []interface{}{ids},
				WithoutParentheses: true,
			},
		}).Limit(limit).Offset(offset).
		Find(&userInfo).
		Error
	return
}

// GetFollowerUserInfo 获取粉丝列表
func GetFollowerUserInfo(ids []string) (userInfo []*forum.UserInfo, total int64, err error) {
	err = global.GVA_DB.Table("frm_users").
		Where("user_id in (?)", ids).
		Clauses(clause.OrderBy{
			Expression: clause.Expr{
				SQL:                "FIELD(user_id, ?)",
				Vars:               []interface{}{ids},
				WithoutParentheses: true,
			},
		}).
		Find(&userInfo).
		Count(&total).
		Error
	return
}
