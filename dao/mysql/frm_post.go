package mysql

import (
	"forum/global"
	"forum/model/forum"

	"gorm.io/gorm/clause"
)

// FrmPostCreatePost 插入帖子数据
func FrmPostCreatePost(p *forum.FrmPost) (err error) {
	return global.GVA_DB.Table("frm_posts").
		Create(&p).
		Error
}

// FrmJudgeCommunityIsExist 判断社区是否存在
func FrmJudgeCommunityIsExist(communityId int64) (isExists bool, err error) {
	var total int64
	err = global.GVA_DB.Table("frm_communities").Select("community_id").Where("community_id = ?", communityId).Count(&total).Error
	if total == 0 {
		return true, err
	} else {
		return false, err
	}
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

// FrmGetPostById 根据帖子
func FrmGetPostById(pid int64) (post *forum.FrmPost, err error) {
	err = global.GVA_DB.Table("frm_posts").
		Select("post_id, content, like_num, author_id, title,community_id,created_at").
		Where("post_id = ?", pid).
		Find(&post).Error
	return
}
