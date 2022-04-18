package main

import (
	"flag"
	"go-port/jenkins/util"
	"strings"
)

func main() {
	city := flag.String("city", "", "city")
	to := flag.String("to", "", "to")
	// 解析命令行中的参数
	flag.Parse()
	// 设置本地时间
	util.SetLocalTime()
	if *to == "" {
		println("请传入to")
		return
	}
	// 字符串分割, 使用字符分割出to
	tos := strings.Split(*to, ";")
	// 获取天气
	table, err := util.GetWeather(*city)
	if err != nil {
		println("获取天气失败，", err)
		return
	}
	// 发送邮件
	err = util.SendGoMail(tos, "天气预报", table, nil)
	if err != nil {
		println("发送邮件失败，", err)
		return
	}
}
