package goutils

import "time"

// 获取当前的毫秒值
func GetCurrentMillisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
