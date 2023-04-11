package v1

import (
	"forum/api/v1/forum"
	"forum/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
	ForumApiGroup  forum.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
