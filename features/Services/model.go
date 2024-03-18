package Services

import (
	"context"
	"log"

	"github.com/bxcodec/faker/v3"
	// "go.mongodb.org/mongo-driver/bson/primitive"

	mongodb "main/core/db/monodb"
)

// Services represents the structure of an Services
// Services represents the structure of a service
type Service struct {
	ID             string `json:"services_id" bson:"_id,omitempty"`
	ServicesTitle  string `json:"services_title" example:"Web Development" validate:"required"`
	ServicesBody   string `json:"services_body" example:"We specialize in building responsive and user-friendly websites." validate:"required"`
	ServicesAssets string `json:"services_assets" example:"http://example.com/assets/service.jpg" validate:"required"`
	ServicesType   string `json:"services_type" example:"Development" validate:"required"`
}

var CTX context.Context = context.Background()

// InitData initializes the database collection and generates fake and provided services.
func InitData() {
	// mongodb.InitMongoDB()
	// Define the list of services
	services := []interface{}{
		Service{
			ServicesTitle:  "Web Development",
			ServicesBody:   "We specialize in building responsive and user-friendly websites.",
			ServicesAssets: "http://example.com/image.jpg",
			ServicesType:   "Development",
		},
		// Service{
		// 	ServicesTitle:  "Graphic Design",
		// 	ServicesBody:   "We provide creative and innovative graphic design solutions.",
		// 	ServicesAssets: "http://example.com/assets/design.jpg",
		// 	ServicesType:   "Design",
		// }, Service{
		// 	ServicesTitle:  "Graphic Design",
		// 	ServicesBody:   "We provide creative and innovative graphic design solutions.",
		// 	ServicesAssets: "http://example.com/assets/design.jpg",
		// 	ServicesType:   "Design",
		// }, Service{
		// 	ServicesTitle:  "Graphic Design",
		// 	ServicesBody:   "We provide creative and innovative graphic design solutions.",
		// 	ServicesAssets: "http://example.com/assets/design.jpg",
		// 	ServicesType:   "Design",
		// }, Service{
		// 	ServicesTitle:  "Graphic Design",
		// 	ServicesBody:   "We provide creative and innovative graphic design solutions.",
		// 	ServicesAssets: "http://example.com/assets/design.jpg",
		// 	ServicesType:   "Design",
		// }, Service{
		// 	ServicesTitle:  "Graphic Design",
		// 	ServicesBody:   "We provide creative and innovative graphic design solutions.",
		// 	ServicesAssets: "http://example.com/assets/design.jpg",
		// 	ServicesType:   "Design",
		// },
	}

	// Insert provided services into MongoDB
	_, err := mongodb.DB.Collection("services").InsertMany(CTX, services)

	if err != nil {
		log.Println("Error inserting services:", err)
	}

	// // Generate and insert fake services
	// fakeServices := GenerateFakeServices(20)
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
