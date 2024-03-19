// Package SocialMedia provides structures and functions related to social media details.
package SocialMedia

import (
	"context"
	"log"

	"github.com/bxcodec/faker/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// "go.mongodb.org/mongo-driver/bson/primitive"

	mongodb "main/core/db/monodb"
)

// SocialMediaDetails represents the structure of social media details of an individual
type SocialMedia struct {
	ID          primitive.ObjectID `json:"sm_id" bson:"_id,omitempty"`
	SmFacebook  string             `json:"sm_facebook" example:"ahmed.mady.facebook" validate:"required"`
	SmWhatsapp  string             `json:"sm_whatsapp" example:"123456789" validate:"required"`
	SmGithub    string             `json:"sm_github" example:"ahmed_mady_github" validate:"required"`
	SmLinkedin  string             `json:"sm_linkedin" example:"john-doe-linkedin" validate:"required"`
	SmEmail     string             `json:"sm_email" example:"ahmed.mady@example.com" validate:"required"`
	SmTwitter   string             `json:"sm_twitter" example:"@ahmed_mady_twitter" validate:"required"`
	SmCv        string             `json:"sm_cv" example:"ahmed_mady_cv.pdf" validate:"required"`
	SmInstagram string             `json:"sm_instagram" example:"ahmed_mady_instagram" validate:"required"`
	SmMedium    string             `json:"sm_medium" example:"@ahmed_mady_medium" validate:"required"`
}

var CTX context.Context = context.Background()

// InitData initializes the database collection and generates fake and provided social media details.
func InitData() {
	// mongodb.InitMongoDB()
	// Define the list of social media details
	socialMediaDetails := []interface{}{
		SocialMedia{
			SmWhatsapp:  "https://wa.me/+201555663045?text=from%20my%20website",
			SmFacebook:  "https://www.facebook.com/ahmedeed.eed.7",
			SmGithub:    "https://github.com/ahmed-eid-faried",
			SmLinkedin:  "https://www.linkedin.com/in/ahmed-mady-a16715228/",
			SmEmail:     "https://mail.google.com/mail/u/0/?fs=1&to=ahmed.eid.ac.1.edu@gmail.com&su=FromWebsite&tf=cm",
			SmCv:        "https://drive.google.com/file/d/1q5Vg44gRgH9Er4mCN5lYJRlVpURrBIdY/view?usp=sharing",
			SmInstagram: "https://www.instagram.com/ahmed_eid_ac/",
			SmTwitter:   "https://twitter.com/AHMEDMA65756172/",
			SmMedium:    "https://medium.com/@ahmed.eid.ac.1.edu",
		},
	}
	// SmEmail:     "ahmed.eid.ac.1.edu@gmail.com",

	// Insert provided social media details into MongoDB
	_, err := mongodb.DB.Collection("social_media").InsertMany(CTX, socialMediaDetails)

	if err != nil {
		log.Println("Error inserting social media details:", err)
	}

	// // Generate and insert fake social media details
	// fakeSocialMediaDetails := GenerateFakeSocialMediaDetails(20)
	// _, err = mongodb.DB.Collection("social_media").InsertMany(CTX, fakeSocialMediaDetails)
	// if err != nil {
	// 	log.Println("Error inserting fake social media details:", err)
	// }
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
			SmMedium:    sm.SmMedium,
		})
	}
	return socialMediaDetails
}
