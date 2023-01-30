package request

import (
	"forum-server/model/common/request"
	"forum-server/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
