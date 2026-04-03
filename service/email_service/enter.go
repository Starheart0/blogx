package email_service

import (
	"blogx_server/global"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendRegisterCode(to, code string) error {
	em := global.Config.Email
	subject := fmt.Sprintf("【%s】账号注册", em.SendNickname)
	text := fmt.Sprintf("你正在进行账号注册操作，这是你的验证码%s ，十分钟内有效", code)
	return SendEmail(to, subject, text)
}

func SendResetPwdCode(to, code string) error {
	em := global.Config.Email
	subject := fmt.Sprintf("【%s】密码重置", em.SendNickname)
	text := fmt.Sprintf("你正在进行密码重置操作，这是你的验证码%s ，十分钟内有效", code)
	return SendEmail(to, subject, text)
}

func SendBindEmailCode(to, code string) error {
	em := global.Config.Email
	subject := fmt.Sprintf("【%s】绑定邮箱", em.SendNickname)
	text := fmt.Sprintf("你正在进行绑定邮箱操作，这是你的验证码%s ，十分钟内有效", code)
	return SendEmail(to, subject, text)
}

func SendEmail(to, subject, text string) (err error) {
	em := global.Config.Email
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", em.SendNickname, em.SendEmail)
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(text)
	err = e.Send(fmt.Sprintf("%s:%d", em.Domain, em.Port), smtp.PlainAuth("", em.SendEmail, em.AuthCode, em.Domain))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ac")
	return nil
}
