// Package SocialMedia provides structures and functions related to social media details.
package SocialMedia

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	mongodb "main/core/db/monodb"
	_ "main/docs" // This is required for Swagger to find your documentation
)

// @Summary Get all social media details
// @Description Retrieve all social media details from MongoDB
// @Tags SocialMediaDetails
// @Produce json
// @Success 200 {array} SocialMedia
// @Router /social_media [get]
func GetSocialMediaDetails(c *gin.Context) {
	var socialMediaDetails []SocialMedia
	cursor, err := mongodb.DB.Collection("social_media").Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error querying social media details:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying social media details"})
		return
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &socialMediaDetails); err != nil {
		log.Println("Error scanning social media details:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning social media details"})
		return
	}

	c.JSON(http.StatusOK, socialMediaDetails)
}

// @Summary Get a social media detail by ID
// @Description Retrieve a social media detail by its ID from MongoDB
// @Tags SocialMediaDetails
// @Produce json
// @Param id path string true "Social media detail ID"
// @Success 200 {object} SocialMedia
// @Router /social_media/{id} [get]
func GetSocialMediaDetailsByID(c *gin.Context) {
	id := c.Param("id")
	var socialMediaDetail SocialMedia
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid social media detail ID"})
		return
	}

	err = mongodb.DB.Collection("social_media").FindOne(context.Background(), bson.M{"_id": objID}).Decode(&socialMediaDetail)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Social media detail not found"})
		return
	}

	c.JSON(http.StatusOK, socialMediaDetail)
}

// @Summary Create a new social media detail
// @Description Create a new social media detail in MongoDB
// @Tags SocialMediaDetails
// @Accept json
// @Produce json
// @Param socialMediaDetail body SocialMedia true "Social media detail object"
// @Success 201 {string} string "ID of the created social media detail"
// @Router /social_media [post]
func CreateSocialMediaDetails(c *gin.Context) {
	var socialMediaDetail SocialMedia
	if err := c.ShouldBindJSON(&socialMediaDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := mongodb.DB.Collection("social_media").InsertOne(context.Background(), socialMediaDetail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating social media detail"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}

// @Summary Update a social media detail
// @Description Update an existing social media detail in MongoDB
// @Tags SocialMediaDetails
// @Accept json
// @Produce json
// @Param id path string true "Social media detail ID"
// @Param socialMediaDetail body SocialMedia true "Updated social media detail object"
// @Success 200
// @Router /social_media/{id} [put]
func UpdateSocialMediaDetails(c *gin.Context) {
	id := c.Param("id")
	var socialMediaDetail SocialMedia
	if err := c.ShouldBindJSON(&socialMediaDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid social media detail ID"})
		return
	}

	_, err = mongodb.DB.Collection("social_media").ReplaceOne(context.Background(), bson.M{"_id": objID}, socialMediaDetail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating social media detail"})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Delete a social media detail
// @Description Delete a social media detail by its ID from MongoDB
// @Tags SocialMediaDetails
// @Param id path string true "Social media detail ID"
// @Success 200
// @Router /social_media/{id} [delete]
func DeleteSocialMediaDetails(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid social media detail ID"})
		return
	}

	_, err = mongodb.DB.Collection("social_media").DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting social media detail"})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Delete all social media details
// @Description Delete all social media details from MongoDB
// @Tags SocialMediaDetails
// @Success 204
// @Router /social_media [delete]
func DeleteAllSocialMediaDetails(c *gin.Context) {
	_, err := mongodb.DB.Collection("social_media").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting social media details"})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary Search social media details
// @Description Search social media details in MongoDB based on a keyword
// @Tags SocialMediaDetails
// @Produce json
// @Param keyword query string true "Search keyword"
// @Success 200 {array} SocialMedia
// @Router /social_media/search [get]
func SearchSocialMediaDetails(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search keyword is required"})
		return
	}

	filter := bson.M{
		"$or": []bson.M{
			{"sm_name": bson.M{"$regex": keyword, "$options": "i"}},
			{"sm_desc": bson.M{"$regex": keyword, "$options": "i"}},
		},
	}

	var socialMediaDetails []SocialMedia
	cursor, err := mongodb.DB.Collection("social_media").Find(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying social media details"})
		return
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &socialMediaDetails); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning social media details"})
		return
	}

	c.JSON(http.StatusOK, socialMediaDetails)
}
