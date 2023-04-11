package request

import (
	"forum/model/common/request"
	"forum/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
