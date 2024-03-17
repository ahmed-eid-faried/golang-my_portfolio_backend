// Package SocialMedia provides structures and functions related to social media details.
package SocialMedia

import (
	"context"
	"log"

	"github.com/bxcodec/faker/v3"
	// "go.mongodb.org/mongo-driver/bson/primitive"

	mongodb "main/core/db/monodb"
)

// SocialMediaDetails represents the structure of social media details of an individual
type SocialMedia struct {
	ID          string `json:"sm_id" bson:"_id,omitempty"`
	SmFacebook  string `json:"sm_facebook" example:"ahmed.mady.facebook" validate:"required"`
	SmWhatsapp  string `json:"sm_whatsapp" example:"123456789" validate:"required"`
	SmGithub    string `json:"sm_github" example:"ahmed_mady_github" validate:"required"`
	SmLinkedin  string `json:"sm_linkedin" example:"john-doe-linkedin" validate:"required"`
	SmEmail     string `json:"sm_email" example:"ahmed.mady@example.com" validate:"required"`
	SmTwitter   string `json:"sm_twitter" example:"@ahmed_mady_twitter" validate:"required"`
	SmCv        string `json:"sm_cv" example:"ahmed_mady_cv.pdf" validate:"required"`
	SmInstagram string `json:"sm_instagram" example:"ahmed_mady_instagram" validate:"required"`
}

var CTX context.Context = context.Background()

// InitData initializes the database collection and generates fake and provided social media details.
func InitData() {
	// mongodb.InitMongoDB()
	// Define the list of social media details
	socialMediaDetails := []interface{}{
		SocialMedia{
			SmFacebook:  "ahmed.mady.facebook",
			SmWhatsapp:  "123456789",
			SmGithub:    "ahmed_mady_github",
			SmLinkedin:  "john-doe-linkedin",
			SmEmail:     "ahmed.mady@example.com",
			SmTwitter:   "@ahmed_mady_twitter",
			SmCv:        "ahmed_mady_cv.pdf",
			SmInstagram: "ahmed_mady_instagram",
		},
		SocialMedia{
			SmFacebook:  "jane.doe.facebook",
			SmWhatsapp:  "987654321",
			SmGithub:    "jane_doe_github",
			SmLinkedin:  "jane-doe-linkedin",
			SmEmail:     "jane.doe@example.com",
			SmTwitter:   "@jane_doe_twitter",
			SmCv:        "jane_doe_cv.pdf",
			SmInstagram: "jane_doe_instagram",
		},
	}

	// Insert provided social media details into MongoDB
	_, err := mongodb.DB.Collection("social_media").InsertMany(CTX, socialMediaDetails)

	if err != nil {
		log.Println("Error inserting social media details:", err)
	}

	// Generate and insert fake social media details
	fakeSocialMediaDetails := GenerateFakeSocialMediaDetails(20)
	_, err = mongodb.DB.Collection("social_media").InsertMany(CTX, fakeSocialMediaDetails)
	if err != nil {
		log.Println("Error inserting fake social media details:", err)
	}
}

// GenerateFakeSocialMediaDetails generates fake social media details
func GenerateFakeSocialMediaDetails(numSocialMediaDetails int) []interface{} {
	var socialMediaDetails []interface{}
	for i := 0; i < numSocialMediaDetails; i++ {
		var sm SocialMedia
		err := faker.FakeData(&sm)
		if err != nil {
			log.Println("Error generating fake data for social media details:", err)
			continue
		}
		socialMediaDetails = append(socialMediaDetails, SocialMedia{
			SmFacebook:  sm.SmFacebook,
			SmWhatsapp:  sm.SmWhatsapp,
			SmGithub:    sm.SmGithub,
			SmLinkedin:  sm.SmLinkedin,
			SmEmail:     sm.SmEmail,
			SmTwitter:   sm.SmTwitter,
			SmCv:        sm.SmCv,
			SmInstagram: sm.SmInstagram,
		})
	}
	return socialMediaDetails
}
