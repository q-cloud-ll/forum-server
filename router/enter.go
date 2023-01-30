package router

import (
	"forum-server/router/example"
	"forum-server/router/forum"
	"forum-server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Forum   forum.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
