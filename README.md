# jenkins

## 发邮件
```
go run cmd/mail.go -to 邮箱地址1 邮箱地址2 -sub 邮件主题 -body 邮件内容 -file 文件地址1 文件地址2
```

## 发送天气预报到用户邮箱
```
go run cmd/weather.go -to 邮箱地址1 邮箱地址2
```

## 获取本机信息
```
go run cmd/runtime.go
```
