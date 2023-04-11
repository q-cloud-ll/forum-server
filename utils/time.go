package utils

import "time"

var TimeTemplates = "2006-01-02 15:04:05" //常规类型

func TimeStringToGoTime(tm, format string) time.Time {
	t, err := time.ParseInLocation(format, tm, time.Local)
	if err == nil && !t.IsZero() {
		return t
	}
	return time.Time{}
}
