package forum

import "forum-server/service"

type ApiGroup struct {
	UserApi
	PostApi
	VoteApi
	CommentApi
	QRCodeApi
	CommunityApi
}

var (
	userService      = service.ServiceGroupApp.ForumServiceGroup.UserService
	postService      = service.ServiceGroupApp.ForumServiceGroup.PostService
	voteService      = service.ServiceGroupApp.ForumServiceGroup.VoteService
	commentService   = service.ServiceGroupApp.ForumServiceGroup.CommentService
	qrcodeService    = service.ServiceGroupApp.ForumServiceGroup.QRCodeService
	communityService = service.ServiceGroupApp.ForumServiceGroup.CommunityService
)
