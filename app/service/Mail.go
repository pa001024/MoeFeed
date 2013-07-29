package service

import (
	"log"
	"net/smtp"
	"strings"
	"time"
)

var (
	auth = smtp.PlainAuth("", _MAIL_FROM, _MAIL_PASS, _MAIL_HOST)
	Mail = &MailService{C: make(chan MailInfo)}
)

// 邮件信息
type MailInfo struct {
	Revi    string
	Title   string
	Context string
}

// 邮件发送服务
type MailService struct {
	C chan MailInfo
}

// 异步发送 默认尝试5次
func (this *MailService) SendMail(rcpt, title, context string) {
	go this.SendMailSync(rcpt, title, context, 5)
}

// 同步发送
func (this *MailService) SendMailSync(rcpt, title, context string, maxtries int) error {
	i := 0
	for {
		err := smtp.SendMail(
			_MAIL_SERVER,
			auth,
			_MAIL_FROM,
			[]string{rcpt},
			[]byte(strings.Join([]string{
				"From: " + _MAIL_FROM,
				"To: " + rcpt,
				"Subject: " + title,
				"Content-Type: text/html",
				"",
				context,
			}, "\r\n")))
		if err != nil {
			log.Println(err)
			i++
			if i > maxtries {
				return err
			}
			time.Sleep(time.Second * 20)
			continue
		}
		return nil
	}
}
