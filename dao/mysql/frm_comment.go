package mysql

import (
	"forum-server/global"
	"forum-server/model/forum"
)

// FrmCreateComment 插入帖子数据
func FrmCreateComment(p *forum.FrmPost) (err error) {
	return global.GVA_DB.Table("frm_posts").
		Create(&p).
		Error
}
