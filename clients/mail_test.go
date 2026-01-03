package clients

import (
	"testing"
)

// TestMailClient_SendMail tests sending mail to MailHog
func TestMailClient_SendMail(t *testing.T) {
	// このテストはdocker-composeでMailHogが起動している必要があります
	// docker-compose up -d mailhog
	// Web UI: http://localhost:17008

	client := NewMailClient("localhost", "17007", "", "", "noreply@golang-gin.test")

	to := []string{"test@example.com"}
	subject := "Test Email from golang-gin"
	body := "This is a test email sent from the golang-gin project.\n\nCheck MailHog UI at http://localhost:17008"

	err := client.SendMail(to, subject, body)
	if err != nil {
		t.Skipf("MailHog not available: %v", err)
		return
	}

	t.Log("Successfully sent email to MailHog")
	t.Log("Check http://localhost:17008 to see the email")
}
