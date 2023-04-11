package forum

import "forum/service"

type ApiGroup struct {
	UserApi
	PostApi
	VoteApi
	CommentApi
	WeChatApi
	CommunityApi
	FollowApi
}

var (
	userService      = service.ServiceGroupApp.ForumServiceGroup.UserService
	postService      = service.ServiceGroupApp.ForumServiceGroup.PostService
	voteService      = service.ServiceGroupApp.ForumServiceGroup.VoteService
	commentService   = service.ServiceGroupApp.ForumServiceGroup.CommentService
	wechatService    = service.ServiceGroupApp.ForumServiceGroup.WeChatService
	communityService = service.ServiceGroupApp.ForumServiceGroup.CommunityService
	followService    = service.ServiceGroupApp.ForumServiceGroup.FollowService
)
