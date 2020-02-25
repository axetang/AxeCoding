/************************************************************************************
 **Author:  Axe Tang; Email: axetang@163.com
 **Package: smtp
 **Element: smtp.SendMail
 **Type: func
 ------------------------------------------------------------------------------------
 **Definition:
 func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
 ------------------------------------------------------------------------------------
 **Description:
 SendMail connects to the server at addr, switches to TLS if possible, authenticates
 with the optional mechanism a if possible, and then sends an email from address from,
 to addresses to, with message msg. The addr must include a port, as in
 "mail.example.com:smtp".
 The addresses in the to parameter are the SMTP RCPT addresses.
 The msg parameter should be an RFC 822-style email with headers first, a blank line,
 and then the message body. The lines of msg should be CRLF terminated. The msg
 headers should usually include fields such as "From", "To", "Subject", and "Cc".
 Sending "Bcc" messages is accomplished by including an email address in the to
 parameter but not including it in the msg headers.
 The SendMail function and the net/smtp package are low-level mechanisms and provide
 no support for DKIM signing, MIME attachments (see the mime/multipart package), or
 other mail functionality. Higher-level packages exist outside of the standard library.
 ------------------------------------------------------------------------------------
 **要点总结:
 1. SendMail连接到addr指定的服务器；如果支持会开启TLS；如果支持会使用a认证身份；然后以from为邮件
 源地址发送邮件msg到目标地址to。（可以是多个目标地址：群发）;
 2.
*************************************************************************************/
package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"strings"
)

func encodeRFC2047(String string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}
func main() {
	// Set up authentication information.
	smtpServer := "smtp.163.com"
	auth := smtp.PlainAuth(
		"",
		"axetang@163.com",
		"163Tf721226",
		smtpServer,
	)
	from := mail.Address{"监控中心", "axetang@163.com"}
	to := mail.Address{"收件人", "axetang72@163.com"}
	to1 := mail.Address{"收件人", "felixtang72@163.com"}
	title := "english Story Sent by Axe."
	body := `The Winepress
A short story by Josef Essberger
grapes"You don't have to be French to enjoy a decent red wine," Charles Jousselin de Gruse used to tell his foreign guests whenever he entertained them in Paris. "But you do have to be French to recognize one," he would add with a laugh.
After a lifetime in the French diplomatic corps, the Count de Gruse lived with his wife in an elegant townhouse on Quai Voltaire. He was a likeable man, cultivated of course, with a well-deserved reputation as a generous host and an amusing raconteur.
`
	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = encodeRFC2047(title)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))
	println(message)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	println("--------Start SendMail--------")
	err := smtp.SendMail(
		smtpServer+":25",
		auth,
		from.Address,
		[]string{to.Address, to1.Address},
		[]byte(message),
	)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Success in Sending mail to axetang72@163.com")
	}
}
