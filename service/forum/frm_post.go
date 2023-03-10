package forum

import (
	"fmt"
	"forum-server/dao/mysql"
	"forum-server/dao/redis"
	"forum-server/global"
	"forum-server/model/forum"
	frmReq "forum-server/model/forum/request"
	frmResp "forum-server/model/forum/response"
	"forum-server/utils"

	"go.uber.org/zap"

	"github.com/pkg/errors"
)

type PostService struct{}

// FrmPostCreatePost 创建帖子服务
func (ps *PostService) FrmPostCreatePost(p *forum.FrmPost) (err error) {
	// 获取雪花算法的uid
	p.PostId = utils.GenID()
	// 将发表的帖子内容保存进数据库
	err = mysql.FrmPostCreatePost(p)
	if err != nil {
		return err
	}
	// 将帖子ID和帖子类型保存进redis，后续取帖子列表用redis的数据结构
	err = redis.FrmPostCreatePost(p.PostId, p.CommunityId)
	return
}

// FrmPostGetPostList 获取帖子列表服务
func (ps *PostService) FrmPostGetPostList(p *frmReq.PostList) (data []*frmResp.FrmPostDetail, err error) {
	if p.CommunityID == 0 {
		data, err = GetPostList(p)
	} else {
		data, err = GetCommunityPostList(p)
	}
	if err != nil {
		global.GVA_LOG.Info("GetPostListService failed", zap.Error(err))
		return nil, err
	}
	return
}

// GetPostList 查询所有帖子并且获取帖子列表
func GetPostList(p *frmReq.PostList) (data []*frmResp.FrmPostDetail, err error) {
	// 根据排序获取帖子(时间/分数) id
	ids, err := redis.FrmGetPostIdsInOrder(p)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return nil, errors.Wrapf(err, "redis.FrmGetPostIdsInOrder(p) return 0 data")
	}

	// 根据id获取帖子数据
	posts, err := mysql.FrmGetPostListByIds(ids)
	for _, post := range posts {
		fmt.Println(post.PostId)
	}
	if err != nil {
		return nil, err
	}
	// 根据id获取每篇帖子的投票数
	voteData, err := redis.FrmGetPostVoteData(ids)
	if err != nil {
		return nil, err
	}

	// 获取作者数据和社区数据，组装每一篇帖子
	for idx, post := range posts {
		// 根据作者id获取作者信息
		user, err := mysql.FrmGetUserById(post.AuthorId)
		if err != nil {
			global.GVA_LOG.Info("GetUserById(post.AuthorId) failed",
				zap.Any("author_id", post.AuthorId),
				zap.Error(err))
			continue
		}

		// 根据社区id获取社区信息
		community, err := mysql.FrmGetCommunityDetailById(post.CommunityId)
		if err != nil {
			global.GVA_LOG.Info("GetCommunityDetailById(post.CommunityId) failed",
				zap.Int64("community_id", post.CommunityId),
				zap.Error(err))
			continue
		}

		// 将得到的数据组装
		postDetail := &frmResp.FrmPostDetail{
			AuthorName:   user.Nickname,
			VoteNum:      voteData[idx],
			FrmPost:      post,
			FrmCommunity: community,
		}

		data = append(data, postDetail)
	}

	return
}

// GetCommunityPostList 根据社区id获取帖子信息
func GetCommunityPostList(p *frmReq.PostList) (data []*frmResp.FrmPostDetail, err error) {
	// 根据社区id查询该社区下的所有帖子id，按排行或者按时间排序
	ids, err := redis.FrmGetCommunityPostIdsInOrder(p)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		global.GVA_LOG.Info("FrmGetCommunityPostIdsInOrder return 0 data")
		return
	}
	// 根据查询到的帖子ids去数据库查询帖子信息，要按给定的id顺序返回帖子内容
	posts, err := mysql.FrmGetPostListByIds(ids)
	if err != nil {
		return nil, err
	}
	// 根据社区id查询社区的详细信息
	voteData, err := redis.FrmGetPostVoteData(ids)
	if err != nil {
		return nil, err
	}

	// 将帖子的作者及分区信息查询出来填充到帖子中
	for idx, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.FrmGetUserById(post.AuthorId)
		if err != nil {
			global.GVA_LOG.Info("FrmGetUserById(post.AuthorId) failed",
				zap.Any("author_id", post.AuthorId),
				zap.Error(err))
			continue
		}

		// 根据社区id查询社区详细信息
		community, err := mysql.FrmGetCommunityDetailById(post.CommunityId)
		if err != nil {
			global.GVA_LOG.Info("FrmGetCommunityDetailById(post.CommunityId) failed",
				zap.Int64("community_id", post.CommunityId),
				zap.Error(err))
			continue
		}
		// 将数据组装起来
		postDetail := &frmResp.FrmPostDetail{
			AuthorName:   user.Nickname,
			VoteNum:      voteData[idx],
			FrmPost:      post,
			FrmCommunity: community,
		}
		data = append(data, postDetail)
	}

	return
}
