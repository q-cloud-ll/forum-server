package request

import (
	"forum/model/common/request"
	"forum/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
