package request

import (
	"forum/model/common/request"
	"forum/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
