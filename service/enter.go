package service

import (
	"forum-server/service/forum"
	"forum-server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	ForumServiceGroup  forum.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
