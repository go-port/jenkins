package util

import (
	"gopkg.in/gomail.v2"
	"os"
)

const (
	HOST = "smtp.163.com"        // 邮件服务器地址
	PORT = 25                    // 端口
	USER = "zg154220830@163.com" // 发送邮件用户账号
	PWD  = "BKBGETSCDNNQTIAF"    // 授权密码
)

/*
SendGoMail 使用gomail发送邮件
@param []string mailAddress 收件人邮箱
@param string subject 邮件主题
@param string body 邮件内容
@param string attaches 附件内容
@return error
*/
func SendGoMail(mailAddress []string, subject string, body string, attaches []string) error {
	m := gomail.NewMessage()
	nickname := "go-port"
	m.SetHeader("From", nickname+"<"+USER+">")
	// 发送给多个用户
	m.SetHeader("To", mailAddress...)
	// 设置邮件主题
	m.SetHeader("Subject", subject)
	// 设置邮件正文
	m.SetBody("text/html", body)
	for _, file := range attaches {
		_, err := os.Stat(file)
		if err != nil {
			println("Error:", file, "does not exist")
		} else {
			println("uploading", file, "...")
			m.Attach(file)
		}
	}
	d := gomail.NewDialer(HOST, PORT, USER, PWD)
	// 发送邮件
	err := d.DialAndSend(m)
	return err
}
