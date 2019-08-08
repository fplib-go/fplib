package fplib

import (
	"time"
)

// "2006-01-02 15:04:05 -0700 MST"
func Datetime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
func Timestamp() int64 {
	return time.Now().Unix()
}
func Timestamp_ms() int64 {
	return int64(Timestamp_ns() / 1e6)
}
func Timestamp_ns() int64 {
	return int64(time.Now().UnixNano())
}
func Datetime_ms() string {
	return time.Now().Format("2006-01-02 15:04:05.0000000")
}
