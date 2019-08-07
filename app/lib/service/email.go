package service

import (
	"net/smtp"
	"strconv"
)

type EmailConfig struct {
	Username string
	Password string
	Host     string
	Port     int
}

func Example() error {
	smtpHost := "localhost"
	smtpPort := 1025
	smtpPass := "yourpassword"
	smtpUser := "yourusername@gmail.com"

	emailConf := &EmailConfig{smtpUser, smtpPass, smtpHost, smtpPort}

	emailauth := smtp.PlainAuth("", emailConf.Username, emailConf.Password, emailConf.Host)

	sender := "fromwho@gmail.com"

	receivers := []string{
		"someone@somedomain.com",
	} // change here

	message := []byte("Hello from Go SMTP package!")

	go func() {
		err := smtp.SendMail(smtpHost+":"+strconv.Itoa(emailConf.Port),
			emailauth,
			sender,
			receivers,
			message,
		)
	}()
	return err
}

