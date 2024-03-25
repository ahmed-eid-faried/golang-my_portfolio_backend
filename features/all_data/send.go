package AllData

import (
	"fmt"
	"net/http"
	"net/smtp"

	"github.com/gin-gonic/gin"
)

// @Summary Send a message
// @Description Send a message
// @ID send-message
// @Accept  json
// @Produce  json
// @Tags Send A Message
// @Param cm_name query string true "Name"
// @Param cm_address query string true "Address"
// @Param cm_number query string true "Phone Number"
// @Param cm_subject query string true "Subject"
// @Param cm_message query string true "Message"
// @Success 200 {string} string "Message sent successfully"
// @Failure 400 {string} string "Bad request"
// @Router /data/messages [post]
func SendMessage(c *gin.Context) {
	name := c.Query("cm_name")
	address := c.Query("cm_address")
	number := c.Query("cm_number")
	subject := c.Query("cm_subject")
	message := c.Query("cm_message")

	if name == "" || address == "" || number == "" || subject == "" || message == "" {
		c.String(http.StatusBadRequest, "Name, Address, Phone Number, Subject, and Message are required")
		return
	}

	// Send email using SMTP
	err := SendEmail(name, address, number, subject, message)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to send message: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, "Message sent successfully")
}

func SendEmail(name, address, number, subject, message string) error {
	// Set up authentication credentials
	auth := smtp.PlainAuth("", "ahmed.eid.ac.1.edu@gmail.com", "A1S2D3F4G5H6J7K8L954321", "smtp.gmail.com")

	// Set up email headers and body
	email := fmt.Sprintf("From: ahmed.eid.ac.1.edu@gmail.com\r\n"+
		"To: ahmed.eid.ac.1.edu@gmail.com, ahmedmadyprof@gmail.com, ahmedmadyprof2@gmail.com\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"Name: %s\r\n"+
		"Address: %s\r\n"+
		"Phone Number: %s\r\n"+
		"Subject: %s\r\n"+
		"Message: %s\r\n", subject, name, address, number, subject, message)

	// Send email
	err := smtp.SendMail("smtp.gmail.com:587", auth, "ahmed.eid.ac.1.edu@gmail.com", []string{"ahmedmadyprof@gmail.com", "ahmedmadyprof2@gmail.com"}, []byte(email))
	if err != nil {
		return err
	}

	return nil
}
