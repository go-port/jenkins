package main

import (
	"go-port/jenkins/util"
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
	println(`运行内存：`, util.FormatByte(mem.TotalAlloc))
}
