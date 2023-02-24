package mysql

import "forum/global"

// FrmPostVoteNum 获取帖子投票数量
func FrmPostVoteNum(postId string) (likeNum int64, err error) {
	var num int64
	err = global.GVA_DB.Table("frm_posts").
		Select("like_num").
		Where("post_id = ?", postId).
		Find(&num).Error
	return num, err
}
