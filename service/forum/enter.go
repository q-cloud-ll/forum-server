package forum

type ServiceGroup struct {
	UserService
	PostService
	VoteService
	CommentService
	WeChatService
	CommunityService
	FollowService
}
