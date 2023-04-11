package router

import (
	"forum/router/forum"
	"forum/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
	Forum  forum.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
