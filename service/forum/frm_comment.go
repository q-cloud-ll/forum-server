package forum

import (
	"errors"
	"forum/dao/mysql"
	"forum/dao/redis"
	"forum/global"
	"forum/model/forum"
	frmReq "forum/model/forum/request"
	"forum/model/forum/response"
	"forum/utils"
	"strconv"

	"go.uber.org/zap"
)

type CommentService struct{}

// FrmCreateComment 新增帖子详情页的评论
func (cs *CommentService) FrmCreateComment(pc *frmReq.PostComment, uid string) (err error) {
	comment := new(forum.FrmComment)
	comment.CommentId = utils.GenID()
	comment.Pid = pc.Pid
	comment.Content = pc.Content
	comment.PostId = pc.PostId
	comment.ReplyId = pc.ReplyId
	comment.UserId = uid
	if pc.Pid != 0 {
		err = redis.FrmIncrChildrenNum(pc.Pid)
		if err != nil {
			global.GVA_LOG.Info("incr子评论数量失败,err:", zap.Int64("parent comment_id:", comment.CommentId), zap.Error(err))
		}
	}
	err = mysql.FrmCreateComment(comment)
	if err != nil {
		return err
	}

	return redis.FrmCreateComment(comment.CommentId, comment.PostId, comment.Pid)
}

// FrmGetPostCommentList 获取帖子评论列表
func (cs *CommentService) FrmGetPostCommentList(pc *frmReq.CommentList) (data []*response.PostComment, total int, err error) {
	// FrmGetCommentsIdsInOrder 根据redis 获取帖子下的根评论id
	ids, err := redis.FrmGetCommentsIdsInOrder(pc)
	if err != nil {
		return nil, 0, err
	}
	if len(ids) == 0 {
		return nil, 0, errors.New("redis.FrmGetCommentIdsInOrder(pc) return 0 data")
	}
	// 如果是根评论，只需要传post_id即可
	// 根据查出来的id查出评论的内容
	comments, err := mysql.FrmGetCommentsListByIds(ids)
	if err != nil {
		return nil, 0, err
	}
	for _, comment := range comments {
		// 获取跟评论的用户信息
		user, err := mysql.FrmGetUserInfoById(comment.UserId)
		if err != nil {
			global.GVA_LOG.Info("FrmGetUserInfoById(comment.UserId) failed",
				zap.String("user_id", comment.UserId),
				zap.Error(err))
			continue
		}
		// 获取每个根评论的子评论数量
		childrenNum, _ := redis.FrmGetChildrenNum(comment.CommentId)

		star, _ := redis.FrmGetCommentStar(strconv.FormatInt(comment.CommentId, 10))

		// 拼接数据
		CommentDetail := &response.PostComment{
			CommentId:   comment.CommentId,
			ChildrenNum: childrenNum,
			CreateTime:  comment.CreatedAt,
			Content:     comment.Content,
			LikeNum:     star,
			PostCreate: &response.PostCreate{
				CreateById:   comment.UserId,
				CreateByName: user.Nickname,
				CreateAvatar: user.Avatar,
			},
			PostReply: &response.PostReply{},
		}
		data = append(data, CommentDetail)
	}
	total = len(data)
	return
}

// FrmGetChildrenCommentList 获取子评论接口
func (cs *CommentService) FrmGetChildrenCommentList(pc *frmReq.CdCommentList) (data []*response.PostComment, total int, err error) {
	// 根据父评论的id获取子评论的id
	ids, err := redis.FrmGetCdCommentsIdsInOrder(pc)
	if err != nil {
		return nil, 0, err
	}
	if len(ids) == 0 {
		return nil, 0, errors.New("redis.FrmGetCommentIdsInOrder(pc) return 0 data")
	}
	// 根据评论id获取评论内容
	comments, err := mysql.FrmGetCommentsListByIds(ids)
	if err != nil {
		return nil, 0, err
	}
	for _, comment := range comments {
		// 获取子评论的用户信息
		user, err := mysql.FrmGetUserInfoById(comment.UserId)
		if err != nil {
			global.GVA_LOG.Info("FrmGetUserInfoById(comment.UserId) failed",
				zap.String("user_id", comment.UserId),
				zap.Error(err))
			continue
		}
		// 获取回复子评论的用户信息
		replyUser, err := mysql.FrmGetUserInfoByCommentId(comment.ReplyId)
		if err != nil {
			global.GVA_LOG.Info("FrmGetUserInfoById(comment.ReplyId) failed",
				zap.Int64("user_id", comment.ReplyId),
				zap.Error(err))
			continue
		}
		// 获取点赞数据
		star, _ := redis.FrmGetCommentStar(strconv.FormatInt(comment.CommentId, 10))

		// 数据拼接
		CommentDetail := &response.PostComment{
			CommentId:  comment.CommentId,
			CreateTime: comment.CreatedAt,
			Content:    comment.Content,
			LikeNum:    star,
			PostCreate: &response.PostCreate{
				CreateById:   comment.UserId,
				CreateByName: user.Nickname,
				CreateAvatar: user.Avatar,
			},
			PostReply: &response.PostReply{
				ReplyUserId: replyUser.UserId,
				ReplyName:   replyUser.Nickname,
				ReplyAvatar: replyUser.Avatar,
			},
		}
		data = append(data, CommentDetail)
	}
	total = len(data)
	return
}

// FrmStarComment 保存评论点赞信息
func (cs *CommentService) FrmStarComment(sc *frmReq.StarComment, uid string) (err error) {
	return redis.FrmStarComment(uid, sc.CommentId, sc.PostId, sc.Pid, float64(sc.Direction))
}
