package service

import (
	"fmt"
	"log"
	"net/smtp"
)

// Mail holds contact form information
type Mail struct {
	// Contact mail address
	from string
	// Mail body
	message string
}

// Sender objects implements Send() method
type Sender interface {
	Send()
}

var _ Sender = &Mail{} // verifies if Mail implements Sender

var hostMail string = "ruidy.nemausat@gmail.com"
var smtpServer string = "smtp.gmail.com:587"

// NewMail creates a Mail with from and body fields
func NewMail(from, body string) Mail {
	return Mail{from, body}
}

// Send runs a mail client and send mail
func (m *Mail) Send() {
	c, err := smtp.Dial(smtpServer)
	if err != nil {
		log.Fatal(err)
	}

	// c.StartTLS()

	if err := c.Mail(hostMail); err != nil {
		log.Fatal(err)
	}

	if err := c.Rcpt(hostMail); err != nil {
		log.Fatal(err)
	}

	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}

	if _, err := fmt.Fprintf(wc, m.message); err != nil {
		log.Fatal(err)
	}

	if err := wc.Close(); err != nil {
		log.Fatal(err)
	}

	err = c.Quit()
	if err != nil {
		log.Fatal(err)
	}
}
