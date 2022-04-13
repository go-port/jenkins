package util

import "time"

// 设置本地时间
func SetLocalTime() {
	// 设置时区为东八
	time.Local = time.FixedZone("CST", 8*3600)
}
