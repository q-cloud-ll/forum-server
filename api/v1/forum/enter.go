package forum

import "forum-server/service"

type ApiGroup struct {
	UserApi
	PostApi
	VoteApi
}

var (
	userService = service.ServiceGroupApp.ForumServiceGroup.UserService
	postService = service.ServiceGroupApp.ForumServiceGroup.PostService
	voteService = service.ServiceGroupApp.ForumServiceGroup.VoteService
)
