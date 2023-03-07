package mail

import (
	"fmt"
	"log"
	"net/smtp"

	"dev.lopo.oma/helper"
	"github.com/jordan-wright/email"
)

var (
	Config   = helper.Load_config()
	sender   = Config.Email.Username
	password = Config.Email.Password
	host     = Config.Email.Host
	port     = Config.Email.Port
	from     = Config.Email.From
	bcc      = Config.Email.Bcc
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

	err := e.Send(Config.Email.Host+":"+fmt.Sprint(Config.Email.Port), smtp.PlainAuth("", sender, password, host))

	if err != nil {
		fmt.Println(err)
		return err
	}

	log.Println("Email sent to " + receiver)
	return nil
}
