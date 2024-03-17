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
)

// @Summary Get all data
// @Description Retrieve Get All Data
// @Tags GetAllData
// @Produce json
// @Success 200 {object} KData
// @Router /data [get]
func GetAllData(c *gin.Context) {
	socialMediaDetails := GetListMongodb(c, "social_media")
	homeDetails := GetListMongodb(c, "home_details")
	projects := GetListMongodb(c, "projects_list")
	services := GetListMongodb(c, "services")

	c.JSON(http.StatusOK, KData{
		SocialMedia:  socialMediaDetails,
		HomeDetails:  homeDetails,
		ProjectsList: projects,
		Services:     services,
	})
}

func GetListMongodb(c *gin.Context, collection string) []interface{} {
	var services []interface{}

	cursor, err4 := mongodb.DB.Collection(collection).Find(context.Background(), bson.M{})
	if err4 != nil {
		log.Printf("Error querying %s: %v\n", collection, err4)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying data"})
		return nil
	}

	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &services); err != nil {
		log.Printf("Error scanning All data: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning All data"})
		return nil
	}
	return services
}

type KData struct {
	SocialMedia  []interface{} `json:"social_media"`
	ProjectsList []interface{} `json:"projects_list"`
	HomeDetails  []interface{} `json:"home_detials"`
	Services     []interface{} `json:"services"`
}
