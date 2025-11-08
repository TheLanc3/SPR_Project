package supplier

import (
	"fmt"
	"net/smtp"
	"spr-project/models"
)

func Supplier(recipient string, position string, quantity int,
	shipmentID int, supp *models.Supplier) {

	from := supp.Email
	password := "cowz fvds sreb cspv" // Use app password for services like Gmail
	to := []string{recipient}
	smtpHost := "smtp.gmail.com" // Example for Gmail
	smtpPort := "587"

	msg := []byte("To: " + recipient + "\r\n" +
		"Subject: The party of product you have requested has been " +
		"shipped with the shipment ID [" +
		fmt.Sprintf("%d", shipmentID) +
		"].\r\n" +
		"\r\n" +
		"Product: " + position + "\r\n" +
		"Quantity: " + fmt.Sprintf("%d", quantity) + "\r\n" +
		"Please indicate the shipment ID [" +
		fmt.Sprintf("%d", shipmentID) +
		"] in the accompanying paper when sending. \r\n\r\n" +
		"Sincerely " + supp.Name + "\r\n" +
		supp.Phone + "\r\n" +
		supp.Address)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}
	fmt.Println("Email from supplier was sent successfully!")
}
