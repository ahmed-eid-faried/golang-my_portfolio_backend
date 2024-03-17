package AllData

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"

	mongodb "main/core/db/monodb"
	_ "main/docs" // This is required for Swagger to find your documentation
	services "main/features/Services"
	home_details "main/features/home_details"
	projects "main/features/projects_list"
	social_media "main/features/social_media"
)

var (
	SocialMedia  []social_media.SocialMedia
	ProjectsList []projects.Project
	HomeDetails  []home_details.HomeDetails
	Services     []services.Service
)

// @Summary Get all data
// @Description Retrieve Get All Data
// @Tags GetAllData
// @Produce json
// @Success 200 {object} KData
// @Router /data [get]
func GetAllData(c *gin.Context) {
	socialMediaDetails := GetListMongodb(c, "social_media", &SocialMedia)
	homeDetails := GetListMongodb(c, "home_details", &HomeDetails)
	projects := GetListMongodb(c, "projects_list", &ProjectsList)
	services := GetListMongodb(c, "services", &Services)

	c.JSON(http.StatusOK, KData{
		SocialMedia:  socialMediaDetails,
		HomeDetails:  homeDetails,
		ProjectsList: projects,
		Services:     services,
	})
}

func GetListMongodb(c *gin.Context, collection string, slicePtr interface{}) interface{} {
	cursor, err := mongodb.DB.Collection(collection).Find(context.Background(), bson.M{})
	if err != nil {
		log.Printf("Error querying %s: %v\n", collection, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return nil
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), slicePtr); err != nil {
		log.Printf("Error scanning All data: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning All data"})
		return nil
	}
	return slicePtr
}

type KData struct {
	SocialMedia  interface{} `json:"social_media"`
	ProjectsList interface{} `json:"projects_list"`
	HomeDetails  interface{} `json:"home_detials"`
	Services     interface{} `json:"services"`
	// SocialMedia  []social_media.SocialMedia `json:"social_media"`
	// ProjectsList []projects.Project         `json:"projects_list"`
	// HomeDetails  []home_details.HomeDetails `json:"home_detials"`
	// Services     []services.Service         `json:"services"`
}
