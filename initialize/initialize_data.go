package initialize

import (
	"forum/global"
	"forum/model/forum"
	"forum/utils"
)

// InitMysqlCommunityData 用于初始化社区数据
func InitMysqlCommunityData() {
	// 创建学习社区、娱乐社区、计算机社区、数学社区
	var communities = []forum.FrmCommunity{
		{CommunityId: utils.GenID(), CommunityName: "学习专区", Introduction: "该社区用于讨论学习考试等，分享资料～"},
		{CommunityId: utils.GenID(), CommunityName: "娱乐专区", Introduction: "该社区可以分享及推荐吃喝玩乐的瞬间～"},
		{CommunityId: utils.GenID(), CommunityName: "计算机专区", Introduction: "该社区可以分享讨论计算机相关问题～"},
		{CommunityId: utils.GenID(), CommunityName: "数学专区", Introduction: "该社区可以讨论数学问题，研究数学理论～"},
	}
	global.GVA_DB.Model(&forum.FrmCommunity{}).Create(&communities)
}
