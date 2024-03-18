package HomeDetails

import (
	"context"
	"log"

	"github.com/bxcodec/faker/v3"
	// "go.mongodb.org/mongo-driver/bson/primitive"

	mongodb "main/core/db/monodb"
)

// HomeDetails represents the structure of an HomeDetails

type HomeDetails struct {
	ID            string `json:"hd_id" bson:"_id,omitempty"`
	HdName        string `json:"hd_name" example:"My Home Details" validate:"required"`
	HdDesc        string `json:"hd_desc" example:"Description of my home details" validate:"required"`
	HdImage       string `json:"hd_image" example:"http://example.com/image.jpg" validate:"required"`
	HdCv          string `json:"hd_cv" example:"http://example.com/cv.pdf" validate:"required"`
	HdAboutmename string `json:"hd_aboutmename" example:"About Me" validate:"required"`
	HdAboutmedesc string `json:"hd_aboutmedesc" example:"I am a software engineer." validate:"required"`
}

var CTX context.Context = context.Background()

// InitData initializes the database collection and generates fake and provided home details.
func InitData() {
	// 	mongodb.InitMongoDB()

	// Define the list of home details
	homeDetails := []interface{}{
		HomeDetails{
			HdName:        "I am a software engineer.",
			HdDesc:        "Passionate Flutter developer with expertise in Dart and Flutter, dedicated to creating high-quality, user-centric mobile apps. Committed to delivering innovative solutions that enhance user experiences and drive business success.",
			HdImage:       "https://i.ibb.co/rxB4Gnm/image.jpg",
			HdCv:          "https://drive.google.com/file/d/1q5Vg44gRgH9Er4mCN5lYJRlVpURrBIdY/view?usp=sharing",
			HdAboutmename: "About Me",
			HdAboutmedesc: "Experienced Flutter Developer skilled in Dart, OOP, and Solid Principles, proficient in problem-solving and data structures. Demonstrated expertise in Flutter, including animations, notifications, and deployment, coupled with strong proficiency in handling RESTful APIs, JSON, Postman, and Thunder Client, along with MySQL and PHP for API integration. Well-versed in state management techniques, including Getx, Provider, and Bloc pattern, with a solid grasp of design patterns, such as MVC, and a commitment to clean code and clean architecture. Proficient in Firebase services, debugging, various testing methodologies, Git, GitHub, UI/UX tools like Adobe XD and Figma, and adept in communication in English. Experienced in CI/CD, Agile development methodologies, and working with Jira for efficient project management.",
		},
	}

	// Insert provided home details into MongoDB
	_, err := mongodb.DB.Collection("home_details").InsertMany(CTX, homeDetails)

	if err != nil {
		log.Println("Error inserting home details:", err)
	}

	// // Generate and insert fake home details
	// fakeHomeDetails := GenerateFakeHomeDetails(20)
	// _, err = mongodb.DB.Collection("home_details").InsertMany(CTX, fakeHomeDetails)
	// if err != nil {
	// 	log.Println("Error inserting fake home details:", err)
	// }
}

// GenerateFakeHomeDetails generates fake home details
func GenerateFakeHomeDetails(numHomeDetails int) []interface{} {
	var homeDetails []interface{}
	for i := 0; i < numHomeDetails; i++ {
		var hd HomeDetails
		err := faker.FakeData(&hd)
		if err != nil {
			log.Println("Error generating fake data for home details:", err)
			continue
		}
		homeDetails = append(homeDetails, HomeDetails{
			HdName:        hd.HdName,
			HdDesc:        hd.HdDesc,
			HdImage:       hd.HdImage,
			HdCv:          hd.HdCv,
			HdAboutmename: hd.HdAboutmename,
			HdAboutmedesc: hd.HdAboutmedesc,
		})
	}
	return homeDetails
}
