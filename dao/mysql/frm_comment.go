package mysql

import (
	"forum/global"
	"forum/model/forum"

	"gorm.io/gorm/clause"
)

// FrmCreateComment 插入帖子数据
func FrmCreateComment(p *forum.FrmComment) (err error) {
	return global.GVA_DB.Table("frm_comments").
		Create(&p).
		Error
}

// FrmGetCommentsListByIds 通过ids查询帖子列表
func FrmGetCommentsListByIds(ids []string) (commentList []*forum.FrmComment, err error) {
	// 要用给定的id顺序返回帖子信息
	err = global.GVA_DB.Table("frm_comments").
		Where("comment_id in (?)", ids).
		Clauses(clause.OrderBy{
			Expression: clause.Expr{
				SQL:                "FIELD(comment_id, ?)",
				Vars:               []interface{}{ids},
				WithoutParentheses: true,
			},
		}).
		Find(&commentList).
		Error
	return
}
