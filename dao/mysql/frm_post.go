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
		Select("user_id, nickname, avatar").
		Where("user_id = ?", uuid).
		Find(&u).
		Error
	return u, err
}

// FrmGetPostById 根据帖子
func FrmGetPostById(pid int64) (post *forum.FrmPost, err error) {
	err = global.GVA_DB.Table("frm_posts").
		Select("post_id, content, like_num, author_id, title,community_id,created_at").
		Where("post_id = ?", pid).
		Find(&post).Error
	return
}
