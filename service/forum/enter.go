package forum

type ServiceGroup struct {
	UserService
	PostService
	VoteService
	CommentService
	QRCodeService
	CommunityService
}
