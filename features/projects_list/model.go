package Projects

import (
	"context"
	"log"

	"github.com/bxcodec/faker/v3"
	// "go.mongodb.org/mongo-driver/bson/primitive"

	mongodb "main/core/db/monodb"
)

// Project represents the structure of a Project
type Project struct {
	ID           string `json:"pl_id" bson:"_id,omitempty"`
	PlTitle      string `json:"pl_title" example:"My Project" validate:"required"`
	PlBody       string `json:"pl_body" example:"Description of my project" validate:"required"`
	PlImage      string `json:"pl_image" example:"" validate:"required"`
	PlGoogleplay string `json:"pl_googleplay" example:"http://example.com/googleplay" validate:"required"`
	PlAppstore   string `json:"pl_appstore" example:"http://example.com/appstore" validate:"required"`
	PlGithub     string `json:"pl_github" example:"http://example.com/github" validate:"required"`
	PlDoc        string `json:"pl_doc" example:"http://example.com/doc" validate:"required"`
	PlPackage    string `json:"pl_package" example:"http://example.com/package" validate:"required"`
	PlCli        string `json:"pl_cli" example:"http://example.com/cli" validate:"required"`
	PlEmbedded   string `json:"pl_embedded" example:"http://example.com/embedded" validate:"required"`
	PlLinux      string `json:"pl_linux" example:"http://example.com/linux" validate:"required"`
	PlWindows    string `json:"pl_windows" example:"http://example.com/windows" validate:"required"`
	PlMacos      string `json:"pl_macos" example:"http://example.com/macos" validate:"required"`
	PlWeb        string `json:"pl_web" example:"http://example.com/web" validate:"required"`
}

var CTX context.Context = context.Background()

// InitData initializes the database collection and generates fake and provided project details.
func InitData() {
	// mongodb.InitMongoDB()

	// Define the list of projects
	projects := []interface{}{
		Project{
			PlTitle:      "PORTFOLIO",
			PlBody:       "A Flutter app designed to showcase professional works and projects, serving as a digital portfolio or resume.",
			PlImage:      "https://i.ibb.co/XYh50Sq/amadytech.png",
			PlGoogleplay: "https://play.google.com/store/apps/details?id=com.amadytech.amadytech",
			PlGithub:     "https://github.com/ahmed-eid-faried/my_portfolio",
			PlDoc:        "https://www.youtube.com/watch?v=ovRCTtGSyV8",
			PlAppstore:   "",
			PlPackage:    "",
			PlCli:        "",
			PlEmbedded:   "",
			PlLinux:      "",
			PlWindows:    "",
			PlMacos:      "",
			PlWeb:        "",
		},
		Project{
			PlTitle:      "Happy Shop",
			PlBody:       "The 'Happy Shop' app offers a fun and easy online shopping experience with features like product search, category browsing, top-sellers, offers, and favorites. Users can receive notifications, track orders, and contact customer service through the app. Developed to enhance user convenience, the Happy Shop app aims to provide an enjoyable shopping experience tailored to users' needs.",
			PlImage:      "https://i.ibb.co/7n0g8tb/happyshop.png",
			PlGoogleplay: "https://play.google.com/store/apps/details?id=com.amadytech.happyshop",
			PlGithub:     "https://github.com/ahmed-eid-faried/happyshop",
			PlDoc:        "https://www.youtube.com/watch?v=RuvCCDwtgtQ&feature=youtu.be",
			PlAppstore:   "",
			PlPackage:    "",
			PlCli:        "",
			PlEmbedded:   "",
			PlLinux:      "",
			PlWindows:    "",
			PlMacos:      "",
			PlWeb:        "",
		},
		Project{
			PlTitle:      "Shafi'I Poems",
			PlBody:       "The 'Imam Shafi'i Poems' application is a great source for exploring and reading the profound religious poems of Imam Shafi'i.",
			PlImage:      "https://i.ibb.co/bH5RwYb/shafii.png",
			PlGoogleplay: "https://play.google.com/store/apps/details?id=com.amadytech.shafii",
			PlGithub:     "https://github.com/ahmed-eid-faried/shafii",
			PlDoc:        "https://www.youtube.com/watch?v=dQrDuOSBGrA",
			PlAppstore:   "",
			PlPackage:    "",
			PlCli:        "",
			PlEmbedded:   "",
			PlLinux:      "",
			PlWindows:    "",
			PlMacos:      "",
			PlWeb:        "",
		},
		Project{
			PlTitle:      "Islamic songs App",
			PlBody:       "The Islamic Songs Application is a distinctive app that offers a collection of high-quality and diverse religious songs. Users can listen to a variety of Islamic songs aimed at enhancing spirituality and contemplation.",
			PlImage:      "https://i.ibb.co/Q9BhgFr/islamicsongs.png",
			PlGoogleplay: "https://play.google.com/store/apps/details?id=com.amadytech.islamic_songs",
			PlGithub:     "https://github.com/ahmed-eid-faried/islamic_songs_app",
			PlDoc:        "https://www.youtube.com/watch?v=rgzgK5Q07lY",
			PlAppstore:   "",
			PlPackage:    "",
			PlCli:        "",
			PlEmbedded:   "",
			PlLinux:      "",
			PlWindows:    "",
			PlMacos:      "",
			PlWeb:        "",
		},
		Project{
			PlTitle:      "MY TASKS",
			PlBody:       "My Tasks is your go-to daily task management app. It's designed to help you stay organized, focused, and in control of your daily activities. With a user-friendly interface and a sleek design, managing your tasks has never been easier.",
			PlImage:      "https://i.ibb.co/RHkT9d3/mytasks.png",
			PlGoogleplay: "https://play.google.com/store/apps/details?id=com.amadytech.mytasks",
			PlGithub:     "https://github.com/ahmed-eid-faried/mytasks",
			PlDoc:        "https://www.youtube.com/watch?v=qkkB0egPwqU",
			PlAppstore:   "",
			PlPackage:    "",
			PlCli:        "",
			PlEmbedded:   "",
			PlLinux:      "",
			PlWindows:    "",
			PlMacos:      "",
			PlWeb:        "",
		},
		Project{
			PlTitle:      "Delivery",
			PlBody:       "Enhance your online shopping experience with our all-in-one Delivery App. Easily manage pending, accepted, and archived orders, receive real-time notifications, contact us via SMS, email, or WhatsApp, track orders on an interactive map with live location, and enjoy smooth animations and reliable Crashlytics integration.",
			PlImage:      "https://i.ibb.co/bzDWGLx/delivery.png",
			PlGoogleplay: "https://play.google.com/store/apps/details?id=com.amadytech.delivery",
			PlGithub:     "https://github.com/ahmed-eid-faried/delivery",
			PlDoc:        "https://www.youtube.com/watch?v=tg2oO5aPWwI",
			PlAppstore:   "",
			PlPackage:    "",
			PlCli:        "",
			PlEmbedded:   "",
			PlLinux:      "",
			PlWindows:    "",
			PlMacos:      "",
			PlWeb:        "",
		},
		Project{
			PlTitle:      "Weather",
			PlBody:       "The Weather App is a useful and essential application available on the Google Play Store that helps you track and know the weather conditions in your area and anywhere else in the world. You can rely on this app to get accurate weather forecasts.",
			PlImage:      "",
			PlGoogleplay: "https://drive.google.com/file/d/1M4ynp3lQs0LIFym2fJt5OwZfKIlagTi1/view?usp=sharing",
			PlGithub:     "https://github.com/ahmed-eid-faried/weather",
			PlDoc:        "https://www.youtube.com/watch?v=4jHMKkqxFQw",
			PlAppstore:   "",
			PlPackage:    "",
			PlCli:        "",
			PlEmbedded:   "",
			PlLinux:      "",
			PlWindows:    "",
			PlMacos:      "",
			PlWeb:        "",
		},
		Project{
			PlTitle:      "News App",
			PlBody:       "The News App is a distinctive application available on the Google Play Store that allows you to stay constantly updated with the	latest news and developments from reliable sources. This app provides you with quick and easy access to the latest local.",
			PlImage:      "https://i.ibb.co/MCQzjSv/news.png",
			PlGoogleplay: "https://drive.google.com/file/d/1JPDMkw3IuEaevIar0Ag30eDtXC-2sX_o/view?usp=sharing",
			PlGithub:     "https://github.com/ahmed-eid-faried/news-app",
			PlDoc:        "https://www.youtube.com/watch?v=Nj2HyHwXijg",
			PlAppstore:   "",
			PlPackage:    "",
			PlCli:        "",
			PlEmbedded:   "",
			PlLinux:      "",
			PlWindows:    "",
			PlMacos:      "",
			PlWeb:        "",
		},
	}

	// Insert provided project details into MongoDB
	result, err := mongodb.DB.Collection("projects_list").InsertMany(CTX, projects)
	if err != nil {
		log.Println("Error inserting projects:", err)
	} else {
		log.Println("Inserted", len(result.InsertedIDs), "projects successfully.")
	}

	// // Generate and insert fake projects
	// fakeprojects := GenerateFakeprojects(20)
	// result, err = mongodb.DB.Collection("projects_list").InsertMany(CTX, fakeprojects)
	// if err != nil {
	// 	log.Println("Error inserting fake projects list:", err)
	// } else {
	// 	log.Println("Inserted", len(result.InsertedIDs), "fake projects successfully.")
	// }
}

// GenerateFakeprojects generates fake projects
func GenerateFakeprojects(numProjects int) []interface{} {
	var projects []interface{}
	for i := 0; i < numProjects; i++ {
		var p Project
		err := faker.FakeData(&p)
		if err != nil {
			log.Println("Error generating fake data for project:", err)
			continue
		}
		projects = append(projects, Project{
			PlTitle:      p.PlTitle,
			PlBody:       p.PlBody,
			PlImage:      p.PlImage,
			PlGoogleplay: p.PlGoogleplay,
			PlAppstore:   p.PlAppstore,
			PlGithub:     p.PlGithub,
			PlDoc:        p.PlDoc,
			PlPackage:    p.PlPackage,
			PlCli:        p.PlCli,
			PlEmbedded:   p.PlEmbedded,
			PlLinux:      p.PlLinux,
			PlWindows:    p.PlWindows,
			PlMacos:      p.PlMacos,
		})
	}
	return projects
}
