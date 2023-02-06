package router

import (
	"forum-server/router/forum"
	"forum-server/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
	Forum  forum.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
