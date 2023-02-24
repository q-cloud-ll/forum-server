package forum

import (
	"forum/dao/mysql"
	"forum/dao/redis"
	"forum/global"
	"forum/model/forum"
	frmReq "forum/model/forum/request"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type FollowService struct{}

// FrmAddFollower 添加关注者（粉丝） followerId 关注者id（粉丝），followeeId 被关注者
func (fs *FollowService) FrmAddFollower(followerId, followeeId string) (err error) {
	err = RedisAddFollowUser(followerId, followeeId)
	if err != nil {
		return err
	}
	err = mysql.AddFollower(followerId, followeeId)
	if err != nil {
		_ = RedisUnFollowUser(followerId, followeeId)
		return err
	}

	return
}

// RedisAddFollowUser 从redis添加
func RedisAddFollowUser(followerId, followeeId string) (err error) {
	// 添加关注者到关注列表
	err = redis.AddFollower(followeeId, followerId)
	if err != nil {
		global.GVA_LOG.Error("AddFollower failed,err:", zap.Error(err))
		return errors.Wrap(err, "AddFollower failed")
	}
	// 添加被关注者到粉丝列表
	err = redis.AddFollowee(followerId, followeeId)
	if err != nil {
		global.GVA_LOG.Error("AddFollowee failed,err:", zap.Error(err))
		// 如果添加失败，则将关注者从关注列表删除
		_ = redis.UnFollower(followeeId, followerId)
		return errors.Wrap(err, "AddFollowee failed")
	}
	return
}

// RedisUnFollowUser redis删除
func RedisUnFollowUser(followerId, followeeId string) (err error) {
	// 从关注列表删除关注者
	err = redis.UnFollower(followeeId, followerId)
	if err != nil {
		return err
	}
	// 从粉丝列表中删除被关注者
	err = redis.UnFollowee(followerId, followeeId)
	if err != nil {
		// 如果删除失败，则将关注者重新添加到关注列表
		_ = redis.AddFollower(followeeId, followerId)
		return err
	}
	return
}

// FrmUnFollower 取关
func (fs *FollowService) FrmUnFollower(followerId, followeeId string) (err error) {
	err = RedisUnFollowUser(followerId, followeeId)
	if err != nil {
		return err
	}
	err = mysql.UnFollowee(followerId, followeeId)
	if err != nil {
		_ = RedisAddFollowUser(followerId, followeeId)
		return err
	}

	return
}

// FrmGetFollowers 获取一个用户的关注列表
func (fs *FollowService) FrmGetFollowers(gf frmReq.GetFollowers) (data []*forum.UserInfo, total int64, err error) {
	followers, err := redis.GetFollowers(gf.UserId)
	if err != nil {
		return nil, 0, err
	}
	data, err = mysql.GetFollowUserInfo(followers, gf.Page, gf.PageSize)
	if err != nil {
		return nil, 0, err
	}
	total = int64(len(data))
	return
}

// FrmGetFollowees 获取一个用户的粉丝列表
func (fs *FollowService) FrmGetFollowees(gf frmReq.GetFollowers) (data []*forum.UserInfo, total int64, err error) {
	followees, err := redis.GetFollowees(gf.UserId, int64(gf.Page), int64(gf.PageSize))
	if err != nil {
		return nil, 0, err
	}
	data, total, err = mysql.GetFollowerUserInfo(followees)
	if err != nil {
		return nil, 0, err
	}

	return
}
