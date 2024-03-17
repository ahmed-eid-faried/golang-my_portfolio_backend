package HomeDetails

import (
	"context"
	"log"
	"net/http"

	// "github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	mongodb "main/core/db/monodb"
	_ "main/docs" // This is required for Swagger to find your documentation
)

// @Summary Get all home details
// @Description Retrieve all home details from MongoDB
// @Tags HomeDetails
// @Produce json
// @Success 200 {array} HomeDetails
// @Router /home_details [get]
func GetHomeDetails(c *gin.Context) {
	var homeDetails []HomeDetails
	cursor, err := mongodb.DB.Collection("home_details").Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error querying home details:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying home details"})
		return
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &homeDetails); err != nil {
		log.Println("Error scanning home details:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning home details"})
		return
	}

	c.JSON(http.StatusOK, homeDetails)
}

// @Summary Get a home detail by ID
// @Description Retrieve a home detail by its ID from MongoDB
// @Tags HomeDetails
// @Produce json
// @Param id path string true "Home detail ID"
// @Success 200 {object} HomeDetails
// @Router /home_details/{id} [get]
func GetHomeDetailsByID(c *gin.Context) {
	id := c.Param("id")
	var homeDetail HomeDetails
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid home detail ID"})
		return
	}

	err = mongodb.DB.Collection("home_details").FindOne(context.Background(), bson.M{"_id": objID}).Decode(&homeDetail)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Home detail not found"})
		return
	}

	c.JSON(http.StatusOK, homeDetail)
}

// @Summary Create a new home detail
// @Description Create a new home detail in MongoDB
// @Tags HomeDetails
// @Accept json
// @Produce json
// @Param homeDetail body HomeDetails true "Home detail object"
// @Success 201 {string} string "ID of the created home detail"
// @Router /home_details [post]
func CreateHomeDetails(c *gin.Context) {
	var homeDetail HomeDetails
	if err := c.ShouldBindJSON(&homeDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := mongodb.DB.Collection("home_details").InsertOne(context.Background(), homeDetail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating home detail"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}

// @Summary Update a home detail
// @Description Update an existing home detail in MongoDB
// @Tags HomeDetails
// @Accept json
// @Produce json
// @Param id path string true "Home detail ID"
// @Param homeDetail body HomeDetails true "Updated home detail object"
// @Success 200
// @Router /home_details/{id} [put]
func UpdateHomeDetails(c *gin.Context) {
	id := c.Param("id")
	var homeDetail HomeDetails
	if err := c.ShouldBindJSON(&homeDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid home detail ID"})
		return
	}

	_, err = mongodb.DB.Collection("home_details").ReplaceOne(context.Background(), bson.M{"_id": objID}, homeDetail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating home detail"})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Delete a home detail
// @Description Delete a home detail by its ID from MongoDB
// @Tags HomeDetails
// @Param id path string true "Home detail ID"
// @Success 200
// @Router /home_details/{id} [delete]
func DeleteHomeDetails(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid home detail ID"})
		return
	}

	_, err = mongodb.DB.Collection("home_details").DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting home detail"})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Delete all home details
// @Description Delete all home details from MongoDB
// @Tags HomeDetails
// @Success 204
// @Router /home_details [delete]
func DeleteAllHomeDetailses(c *gin.Context) {
	_, err := mongodb.DB.Collection("home_details").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting home details"})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary Search home details
// @Description Search home details in MongoDB based on a keyword
// @Tags HomeDetails
// @Produce json
// @Param keyword query string true "Search keyword"
// @Success 200 {array} HomeDetails
// @Router /home_details/search [get]
func SearchHomeDetailses(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search keyword is required"})
		return
	}

	filter := bson.M{
		"$or": []bson.M{
			{"hd_name": bson.M{"$regex": keyword, "$options": "i"}},
			{"hd_desc": bson.M{"$regex": keyword, "$options": "i"}},
		},
	}

	var homeDetails []HomeDetails
	cursor, err := mongodb.DB.Collection("home_details").Find(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying home details"})
		return
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &homeDetails); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning home details"})
		return
	}

	c.JSON(http.StatusOK, homeDetails)
}
