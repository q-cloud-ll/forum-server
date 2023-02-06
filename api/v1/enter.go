package v1

import (
	"forum-server/api/v1/forum"
	"forum-server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
	ForumApiGroup  forum.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
