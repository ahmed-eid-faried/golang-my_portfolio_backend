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
	ID           string    `json:"pl_id" bson:"_id,omitempty"`
	PlTitle      string `json:"pl_title" example:"My Project" validate:"required"`
	PlBody       string `json:"pl_body" example:"Description of my project" validate:"required"`
	PlImage      string `json:"pl_image" example:"http://example.com/image.jpg" validate:"required"`
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
}

var CTX context.Context = context.Background()

// InitData initializes the database collection and generates fake and provided project details.
func InitData() {
	// mongodb.InitMongoDB()

	// Define the list of projects
	projects := []interface{}{
		Project{
			PlTitle:      "My Project",
			PlBody:       "Description of my project",
			PlImage:      "http://example.com/image.jpg",
			PlGoogleplay: "http://example.com/googleplay",
			PlAppstore:   "http://example.com/appstore",
			PlGithub:     "http://example.com/github",
			PlDoc:        "http://example.com/doc",
			PlPackage:    "http://example.com/package",
			PlCli:        "http://example.com/cli",
			PlEmbedded:   "http://example.com/embedded",
			PlLinux:      "http://example.com/linux",
			PlWindows:    "http://example.com/windows",
			PlMacos:      "http://example.com/macos",
		},
		Project{
			PlTitle:      "Another Project",
			PlBody:       "Description of another project",
			PlImage:      "http://example.com/another_image.jpg",
			PlGoogleplay: "http://example.com/another_googleplay",
			PlAppstore:   "http://example.com/another_appstore",
			PlGithub:     "http://example.com/another_github",
			PlDoc:        "http://example.com/another_doc",
			PlPackage:    "http://example.com/another_package",
			PlCli:        "http://example.com/another_cli",
			PlEmbedded:   "http://example.com/another_embedded",
			PlLinux:      "http://example.com/another_linux",
			PlWindows:    "http://example.com/another_windows",
			PlMacos:      "http://example.com/another_macos",
		},
	}

	// Insert provided project details into MongoDB
	result, err := mongodb.DB.Collection("projects_list").InsertMany(CTX, projects)
	if err != nil {
		log.Println("Error inserting projects:", err)
	} else {
		log.Println("Inserted", len(result.InsertedIDs), "projects successfully.")
	}

	// Generate and insert fake projects
	fakeprojects := GenerateFakeprojects(20)
	result, err = mongodb.DB.Collection("projects_list").InsertMany(CTX, fakeprojects)
	if err != nil {
		log.Println("Error inserting fake projects list:", err)
	} else {
		log.Println("Inserted", len(result.InsertedIDs), "fake projects successfully.")
	}
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
