package request

import "forum/model/common/request"

type AddFollower struct {
	FolloweeId string `json:"followee_id"`
}

type GetFollowers struct {
	UserId string `json:"user_id"`
	request.PageInfo
}
