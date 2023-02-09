package forum

type RouterGroup struct {
	UserRouter
	PostRouter
	VoteRouter
	CommentRouter
	QRCodeRouter
	CommunityRouter
}
