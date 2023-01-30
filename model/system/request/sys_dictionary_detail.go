package request

import (
	"forum-server/model/common/request"
	"forum-server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
