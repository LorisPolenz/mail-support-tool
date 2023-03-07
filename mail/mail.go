package mail

import (
	"fmt"
	"log"
	"net/smtp"

	"dev.lopo.mail-support/helper"
	"github.com/jordan-wright/email"
)

var (
	sender   = helper.Configuration.Email.Username
	password = helper.Configuration.Email.Password
	host     = helper.Configuration.Email.Host
	port     = helper.Configuration.Email.Port
	from     = helper.Configuration.Email.From
	bcc      = helper.Configuration.Email.Bcc
)

func append_attachment(path string, e *email.Email) *email.Email {
	file_paths := helper.Extract_file_paths(path)

	for _, file_path := range file_paths {
		_, err := e.AttachFile(path + file_path)
		if err != nil {
			log.Fatal(err)
		}
	}

	return e
}

func Send_mail(path string, receiver string) error {
	e := email.NewEmail()
	e.From = from
	e.To = []string{receiver}
	e.Bcc = []string{bcc}
	e.Subject = GenerateMessageSubject()
	e.HTML = []byte(GenerateMessageBody())

	e = append_attachment(path, e)

	fmt.Println("Sending email to " + host + ":" + fmt.Sprint(port))

	decryptedPassword, _ := helper.Decrypt(password)

	err := e.Send(host+":"+fmt.Sprint(port), smtp.PlainAuth("", sender, decryptedPassword, host))

	if err != nil {
		fmt.Println(err)
		return err
	}

	log.Println("Email sent to " + receiver)
	return nil
}
