package rsacrypt

import (
	"fmt"
	"log"
	"os/user"

	"github.com/scorredoira/email"
)

// Send email to user @verisign.com along with attached files
func SendEmail(uname string, attachfile ...string) error {
	toUserEmail := uname + "@verisign.com"
	log.Println("Sending email to ", toUserEmail)
	config := ReadConfig()
	mailserver := config.Mailserver
	fromUserEmail := config.FromEmail
	subject := "Encrypted file from " + username()
	body := "Attached is encryped file from " + username() + "\r\n" +
		"To decrypt run 'go-rsacrypt -mode=decrypt -keyfile=<path to your private key> -in=<path to encrypted file> -out=<target path to decrypt file>"

	m := email.NewMessage(subject, body)
	m.From = fromUserEmail
	m.To = []string{toUserEmail}
	for _, f := range attachfile {
		err := m.Attach(f)
		if err != nil {
			return fmt.Errorf("Unable to attach file to email. %s", err)
		}
	}

	err := email.Send(mailserver, nil, m)
	if err != nil {
		return fmt.Errorf("Unable to send email. %s", err)
	}

	log.Println("Sending email ..done")
	return err

}

// Return username from os
func username() string {
	u, err := user.Current()
	if err != nil {
		log.Fatalf("Current: %v", err)
	}
	return u.Username
}
