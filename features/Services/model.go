package Services

import (
	"context"
	"log"

	"github.com/bxcodec/faker/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// "go.mongodb.org/mongo-driver/bson/primitive"

	mongodb "main/core/db/monodb"
)

// Services represents the structure of an Services
// Services represents the structure of a service
type Service struct {
	ID             primitive.ObjectID `json:"services_id" bson:"_id,omitempty"`
	ServicesTitle  string             `json:"services_title" example:"Web Development" validate:"required"`
	ServicesBody   string             `json:"services_body" example:"We specialize in building responsive and user-friendly websites." validate:"required"`
	ServicesAssets string             `json:"services_assets" example:"http://example.com/assets/service.jpg" validate:"required"`
	ServicesType   string             `json:"services_type" example:"Development" validate:"required"`
}

var CTX context.Context = context.Background()

// InitData initializes the database collection and generates fake and provided services.
func InitData() {
	// mongodb.InitMongoDB()
	// Define the list of services
	services := []interface{}{
		Service{
			ServicesTitle:  "Web Development",
			ServicesBody:   "We specialize in building responsive and user-friendly websites using React or Flutter.",
			ServicesAssets: "https://i.ibb.co/CVndDjn/web.png",
			ServicesType:   "Web Development",
		},
		Service{
			ServicesTitle:  "Mobile App Development (Android)",
			ServicesBody:   "We develop native Android applications using Flutter or Golang.",
			ServicesAssets: "https://i.ibb.co/j4H29rK/icons8-google-play-100.png",
			ServicesType:   "Mobile Development",
		},
		Service{
			ServicesTitle:  "Mobile App Development (iOS)",
			ServicesBody:   "We develop native iOS applications using Flutter and Golang.",
			ServicesAssets: "https://i.ibb.co/JCstCQH/appstore.png",
			ServicesType:   "Mobile Development",
		},
		Service{
			ServicesTitle:  "Cross-Platform App Development",
			ServicesBody:   "We create cross-platform apps using Flutter and Golang.",
			ServicesAssets: "https://i.ibb.co/7tD4hdX/icons8-multiple-devices-100.png",
			ServicesType:   "Cross-Platform Development",
		},
		Service{
			ServicesTitle:  "Desktop App Development (Windows)",
			ServicesBody:   "We build desktop applications for Windows using Flutter or other technologies.",
			ServicesAssets: "https://i.ibb.co/BZScnXb/icons8-windows-10-100.png",
			ServicesType:   "Desktop Development",
		},
		Service{
			ServicesTitle:  "Desktop App Development (Mac OS)",
			ServicesBody:   "We build desktop applications for Mac OS using Flutter or other technologies.",
			ServicesAssets: "https://i.ibb.co/nzmQ4FL/icons8-mac-client-64.png",
			ServicesType:   "Desktop Development",
		},
		Service{
			ServicesTitle:  "Desktop App Development (Linux)",
			ServicesBody:   "We build desktop applications for Linux using Flutter or other technologies.",
			ServicesAssets: "https://i.ibb.co/Q6GbLsw/icons8-linux-server-100-1.png",
			ServicesType:   "Desktop Development",
		},
		Service{
			ServicesTitle:  "Backend Development",
			ServicesBody:   "We create robust and scalable backend solutions using Golang.",
			ServicesAssets: "https://i.ibb.co/NYmR5Hf/icons8-backend-64.png",
			ServicesType:   "Backend Development",
		},
		Service{
			ServicesTitle:  "Package Tool Development",
			ServicesBody:   "We develop custom package tools to streamline your development process.",
			ServicesAssets: "https://i.ibb.co/crrsbgg/share.png",
			ServicesType:   "Package Tool Development",
		},
		Service{
			ServicesTitle:  "Embedded Systems Development",
			ServicesBody:   "We design and develop embedded systems solutions using Golang or Flutter.",
			ServicesAssets: "https://i.ibb.co/68rq4gN/icons8-embedded-66-1.png",
			ServicesType:   "Embedded Systems Development",
		},
		Service{
			ServicesTitle:  "CLI (Command-Line Interface) Development",
			ServicesBody:   "We build command-line interfaces using Golang for various purposes.",
			ServicesAssets: "https://i.ibb.co/Q8m1H46/coding.png",
			ServicesType:   "CLI Development",
		},
	}
	mongodb.Index(mongodb.DB.Collection("services"))

	// Insert provided services into MongoDB
	_, err := mongodb.DB.Collection("services").InsertMany(CTX, services)

	if err != nil {
		log.Println("Error inserting services:", err)
	}

	// // Generate and insert fake services
	// fakeServices := GenerateFakeServices(20)
	// 	mongodb.Index(mongodb.DB.Collection("services"))
	// _, err = mongodb.DB.Collection("services").InsertMany(CTX, fakeServices)
	// if err != nil {
	// 	log.Println("Error inserting fake services:", err)
	// }
}

// GenerateFakeServices generates fake services
func GenerateFakeServices(numServices int) []interface{} {
	var services []interface{}
	for i := 0; i < numServices; i++ {
		var svc Service
		err := faker.FakeData(&svc)
		if err != nil {
			log.Println("Error generating fake data for services:", err)
			continue
		}
		services = append(services, Service{
			ServicesTitle:  svc.ServicesTitle,
			ServicesBody:   svc.ServicesBody,
			ServicesAssets: svc.ServicesAssets,
			ServicesType:   svc.ServicesType,
		})
	}
	return services
}
