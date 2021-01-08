package tools

import (
	"emanager/pkg/config"
	"net/smtp"
	"strings"
)

/* SendMail
 *  to: example@example.com;example1@163.com;example2@sina.com.cn;...
 *  subject:The subject of mail
 *  body: The content of mail
 */
func SendMail(to string, subject string, body string) error {
	user := config.Config.EmailNotification.Username
	password := config.Config.EmailNotification.Password
	host := config.Config.EmailNotification.Host

	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var contentType string
	contentType = "Content-type:text/html;charset=utf-8"

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	return err
}
