package handlers

import (
	_ "encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/smtp"
	"os"
)

type EmailInput struct {
	From    string `json:"from"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

// Email godoc
// @Summary Send Email
// @Description Sends an email using the data from a form
// @Tags email
// @Accept  json
// @Produce  json
// @Param   email     body    EmailInput   true "Email details"
// @Success 200 {object} EmailInput
// @Router /api/email [post]
func Email(ctx *fiber.Ctx) error {
	// Parse request
	input := new(EmailInput)
	if err := ctx.BodyParser(input); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Set the receiver email
	toEmail := "agfelgue@gmail.com"

	// Send the email
	err := sendEmail(toEmail, input)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Failed to send the email"})
	}

	// Response
	return ctx.JSON(fiber.Map{"version": "1.0"})
}

func sendEmail(toEmail string, input *EmailInput) error {
	from := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	subject := "Email from website"

	body := fmt.Sprintf(`<p><b>%s</b> just sent an email from the website! their message:</p>
	
	<p><b>Subject:</b> %s</p>
	
	<p><b>Message:</b><br/>
	%s
	</p>
	`, input.From, input.Subject, input.Message)

	// Message
	message := []byte("MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		"To: " + toEmail + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{toEmail}, message)
	if err != nil {
		return err
	}
	return nil
}
