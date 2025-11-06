package mailing

import (
	"fmt"
	"net/smtp"
)

func SendEmail(recipient string, position string, quantity int,
	shipmentID int) {

	from := "spr.project2025@gmail.com"
	password := "obko epnp bgxa piol" // Use app password for services like Gmail
	to := []string{recipient}
	smtpHost := "smtp.gmail.com" // Example for Gmail
	smtpPort := "587"

	msg := []byte("To: " + recipient + "\r\n" +
		"Subject: Ship a new party of product.\r\n" +
		"\r\n" +
		"Order: " + position + "\r\n" +
		"Quantity: " + fmt.Sprintf("%d", quantity) + "\r\n" +
		"Please indicate the shipment ID [" +
		fmt.Sprintf("%d", shipmentID) +
		"] in the accompanying paper when sending.")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}
	fmt.Println("Email sent successfully!")
}
