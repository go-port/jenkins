package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	println(`电脑名称：`, name)
	println(`系统类型：`, runtime.GOOS)
	println(`系统架构：`, runtime.GOARCH)
	println(`CPU核数：`, runtime.GOMAXPROCS(0))
	println(`go版本：`, runtime.Version())
	println(`GOROOT：`, runtime.GOROOT())
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	println(`运行内存：`, formatSize(mem.TotalAlloc))
}

// 字节的单位转换 保留两位小数
func formatSize(size uint64) string {
	if size < 1024 {
		//return strconv.FormatInt(size, 10) + "B"
		return fmt.Sprintf("%.2fB", float64(size)/float64(1))
	} else if size < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(size)/float64(1024))
	} else if size < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(size)/float64(1024*1024))
	} else if size < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(size)/float64(1024*1024*1024))
	} else if size < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(size)/float64(1024*1024*1024*1024))
	} else if size < (1024 * 1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fEB", float64(size)/float64(1024*1024*1024*1024*1024))
	} else {
		return fmt.Sprintf("%.2fB", float64(size)/float64(1))
	}
}
