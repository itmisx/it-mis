package mailx

import (
	"it-mis/init/config"

	"gopkg.in/gomail.v2"
)

// SendEmail 发送邮件
// from 发送方
// to 接收方
// suject 主题
// body 内容
func SendEmail(
	from,
	to,
	subject,
	body string,
) error {
	dialer := gomail.NewDialer(
		config.Config.Email.Host,
		config.Config.Email.Port,
		config.Config.Email.Username,
		config.Config.Email.Password,
	)
	dialer.SSL = true
	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)
	return dialer.DialAndSend(msg)
}
