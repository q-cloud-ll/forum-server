package service

import (
	"forum/service/forum"
	"forum/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	ForumServiceGroup  forum.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
