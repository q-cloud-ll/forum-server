package service

import (
	"forum-server/service/example"
	"forum-server/service/forum"
	"forum-server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	ForumServiceGroup   forum.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
