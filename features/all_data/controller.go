package AllData

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"

	core "main/core"
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
// @Tags AllData
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

// @Summary Put All Data
// @Description init all Data from MongoDB
// @Tags AllData
// @Success 200
// @Router /data/init [get]
func PostAllData(c *gin.Context) {
	core.InitDataBase()
	c.Status(http.StatusNoContent)
}

// @Summary Delete all Data
// @Description Delete all Data from MongoDB
// @Tags AllData
// @Success 204
// @Router /data [delete]
func DeleteAllData(c *gin.Context) {
	////////////////////////************home details*************\\\\\\\\\\\\\\\\\\\\\\\\\\
	_, err := mongodb.DB.Collection("home_details").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting home details"})
		return
	}
	////////////////////////*************projects************\\\\\\\\\\\\\\\\\\\\\\\\\\	////////////////////////*************************\\\\\\\\\\\\\\\\\\\\\\\\\\
	_, err = mongodb.DB.Collection("projects_list").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting projects"})
		return
	}
	////////////////////////************services*************\\\\\\\\\\\\\\\\\\\\\\\\\\	////////////////////////*************************\\\\\\\\\\\\\\\\\\\\\\\\\\
	_, err = mongodb.DB.Collection("services").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting services"})
		return
	}
	////////////////////////*************social media details************\\\\\\\\\\\\\\\\\\\\\\\\\\	////////////////////////*************************\\\\\\\\\\\\\\\\\\\\\\\\\\
	_, err = mongodb.DB.Collection("social_media").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting social media details"})
		return
	}
	////////////////////////*************************\\\\\\\\\\\\\\\\\\\\\\\\\\

	c.Status(http.StatusNoContent)
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
