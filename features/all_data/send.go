package AllData

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
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

// func SendEmail(name, address, number, subject, message string) error {
// 	// Set up authentication credentials
// 	auth := smtp.PlainAuth("", "ahmed.eid.ac.1.edu@gmail.com", "A1S2D3F4G5H6J7K8L954321", "smtp.gmail.com")

// 	// Set up email headers and body
// 	email := fmt.Sprintf("From: ahmed.eid.ac.1.edu@gmail.com\r\n"+
// 		"To: ahmed.eid.ac.1.edu@gmail.com, ahmedmadyprof@gmail.com, ahmedmadyprof2@gmail.com\r\n"+
// 		"Subject: %s\r\n"+
// 		"\r\n"+
// 		"Name: %s\r\n"+
// 		"Address: %s\r\n"+
// 		"Phone Number: %s\r\n"+
// 		"Subject: %s\r\n"+
// 		"Message: %s\r\n", subject, name, address, number, subject, message)

// 	// Send email
// 	err := smtp.SendMail("smtp.gmail.com:587", auth, "ahmed.eid.ac.1.edu@gmail.com", []string{"ahmedmadyprof@gmail.com", "ahmedmadyprof2@gmail.com"}, []byte(email))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// Function to send email using Gmail SMTP server and OAuth2 authentication
func SendEmail(name, address, number, subject, message string) error {
	// Create OAuth2 token from client credentials
	ctx := context.Background()
	config, err := google.ConfigFromJSON([]byte(`{
		"web": {
			"client_id": "508930024390-jiusp5cjae7750t731chcn31oilkcl7l.apps.googleusercontent.com",
			"client_secret": "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",
			"redirect_uris": ["https://github.com/ahmed-eid-faried/ahmed-eid-faried.github.io"],
			"auth_uri": "https://accounts.google.com/o/oauth2/auth",
			"token_uri": "https://oauth2.googleapis.com/token",
			"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
			"project_id": "myportfolio-417512"
		}
	}`), gmail.GmailSendScope)
	if err != nil {
		return fmt.Errorf("error creating OAuth2 config: %v", err)
	}
	token := &oauth2.Token{
		AccessToken:  config.Endpoint.TokenURL, // Replace with your actual access token
		RefreshToken: config.Endpoint.TokenURL, // Replace with your actual refresh token
		TokenType:    "Bearer",
	}
	client := config.Client(ctx, token)

	// Create Gmail service
	srv, err := gmail.New(client)
	if err != nil {
		return fmt.Errorf("error creating Gmail service: %v", err)
	}

	// Construct email message
	var msg gmail.Message
	msgStr := []byte(fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"Name: %s\r\n"+
		"Address: %s\r\n"+
		"Phone Number: %s\r\n"+
		"Subject: %s\r\n"+
		"Message: %s\r\n", "ahmed.eid.ac.1.edu@gmail.com", "ahmedmadyprof@gmail.com, ahmedmadyprof2@gmail.com", subject, name, address, number, subject, message))
	msg.Raw = base64.URLEncoding.EncodeToString(msgStr)

	// Send email
	_, err = srv.Users.Messages.Send("me", &msg).Do()
	if err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}

	return nil
}
