package clients

import (
	"fmt"
	"net/smtp"
)

// MailClient wraps SMTP client for sending emails
type MailClient struct {
	host     string
	port     string
	username string
	password string
	from     string
}

// NewMailClient creates a new mail client
func NewMailClient(host, port, username, password, from string) *MailClient {
	return &MailClient{
		host:     host,
		port:     port,
		username: username,
		password: password,
		from:     from,
	}
}

// SendMail sends an email
func (c *MailClient) SendMail(to []string, subject, body string) error {
	message := []byte(fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", c.from, to[0], subject, body))

	addr := fmt.Sprintf("%s:%s", c.host, c.port)

	// For MailHog, no authentication needed
	if c.host == "mailhog" || c.host == "localhost" {
		return smtp.SendMail(addr, nil, c.from, to, message)
	}

	// For real SMTP servers with authentication
	auth := smtp.PlainAuth("", c.username, c.password, c.host)
	return smtp.SendMail(addr, auth, c.from, to, message)
}
