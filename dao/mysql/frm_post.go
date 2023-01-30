package mysql

import (
	"forum-server/global"
	"forum-server/model/forum"

	"gorm.io/gorm/clause"

	uuid "github.com/satori/go.uuid"
)

// FrmPostCreatePost 插入帖子数据
func FrmPostCreatePost(p *forum.FrmPost) (err error) {
	return global.GVA_DB.Table("frm_posts").
		Create(&p).
		Error
}

// FrmGetPostListByIds 通过ids查询帖子列表
func FrmGetPostListByIds(ids []string) (postList []*forum.FrmPost, err error) {
	var list []*forum.FrmPost
	// 要用给定的id顺序返回帖子信息
	err = global.GVA_DB.Table("frm_posts").
		Where("post_id in (?)", ids).
		Clauses(clause.OrderBy{
			Expression: clause.Expr{
				SQL:                "FIELD(post_id, ?)",
				Vars:               []interface{}{ids},
				WithoutParentheses: true,
			},
		}).
		Find(&list).
		Error
	return list, err
}

// FrmGetUserById 根据uuid查询用户信息
func FrmGetUserById(uuid uuid.UUID) (user *forum.FrmUser, err error) {
	var u *forum.FrmUser
	err = global.GVA_DB.Table("frm_users").
		Select("user_id, nickname").
		Where("user_id = ?", uuid).
		Find(&u).
		Error
	return u, err
}
