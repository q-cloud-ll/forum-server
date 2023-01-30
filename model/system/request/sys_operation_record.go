package request

import (
	"forum-server/model/common/request"
	"forum-server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
