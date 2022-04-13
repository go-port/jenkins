package main

import (
	"flag"
	"go-port/jenkins/util"
	"strings"
)

func main() {
	to := flag.String("to", "", "to")        // 接受邮箱
	sub := flag.String("sub", "", "subject") // 邮件主题
	body := flag.String("body", "", "body")  // 邮件内容
	file := flag.String("file", "", "file")  // 邮件附件地址
	// 解析命令行中的参数
	flag.Parse()
	// 设置本地时间
	util.SetLocalTime()
	if *to == "" {
		println("请传入收件人邮箱(to)")
		return
	}
	// 字符串分割, 使用字符分割出to,file
	tos := strings.Split(*to, ";")
	files := make([]string, 0)
	if *file != "" {
		files = strings.Split(*file, ";")
	}
	// 发送邮件
	err := util.SendGoMail(tos, *sub, *body, files)
	if err != nil {
		println("发送邮件失败，", err)
		return
	}
}
