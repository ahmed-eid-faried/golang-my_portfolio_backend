package Services

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	mongodb "main/core/db/monodb"
	_ "main/docs" // This is required for Swagger to find your documentation
)

// GetServices retrieves all services from MongoDB
// @Summary Get All Services
// @Description Get all Services
// @Produce json
// @Success 200 {array} Service
// @Router /services [get]
// @Tags Services
func GetServices(c *gin.Context) {
	var services []Service
	cursor, err := mongodb.DB.Collection("services").Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error querying services:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying services"})
		return
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &services); err != nil {
		log.Println("Error scanning services:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning services"})
		return
	}

	c.JSON(http.StatusOK, services)
}

// GetServicesByID retrieves a service by its ID from MongoDB
// @Summary Get a Service by ID
// @Description Get a Service by ID
// @Produce json
// @Param id path int true "Service ID"
// @Success 200 {object} Service
// @Router /services/{id} [get]
// @Tags Services
func GetServicesByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}

	var service Service
	err = mongodb.DB.Collection("services").FindOne(context.Background(), bson.M{"_id": id}).Decode(&service)
	if err != nil {
		// if err == mongo.ErrNoDocuments {
		// 	c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		// 	return
		// }
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, service)
}

// CreateServices creates a new service in MongoDB
// @Summary Create a Service
// @Description Create a Service
// @Accept json
// @Produce json
// @Param service body Service true "Service object"
// @Success 201
// @Router /services [post]
// @Tags Services
func CreateServices(c *gin.Context) {
	var service Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := mongodb.DB.Collection("services").InsertOne(context.Background(), service)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating service"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}

// UpdateServices updates an existing service in MongoDB
// @Summary Update a Service
// @Description Update a Service by ID
// @Accept json
// @Produce json
// @Param id path int true "Service ID"
// @Param service body Service true "Service object"
// @Success 200
// @Router /services/{id} [put]
// @Tags Services
func UpdateServices(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}

	var service Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = mongodb.DB.Collection("services").ReplaceOne(context.Background(), bson.M{"_id": id}, service)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating service"})
		return
	}

	c.Status(http.StatusOK)
}

// DeleteServices deletes a service by its ID from MongoDB
// @Summary Delete a Service
// @Description Delete a Service by ID
// @Param id path int true "Service ID"
// @Success 200
// @Router /services/{id} [delete]
// @Tags Services
func DeleteServices(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}

	_, err = mongodb.DB.Collection("services").DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting service"})
		return
	}

	c.Status(http.StatusOK)
}

// DeleteAllServices deletes all services from MongoDB
// @Summary Remove all Services
// @Description Remove all Services from the database
// @Produce json
// @Success 204
// @Router /services [delete]
// @Tags Services
func DeleteAllServices(c *gin.Context) {
	_, err := mongodb.DB.Collection("services").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting services"})
		return
	}

	c.Status(http.StatusNoContent)
}

// SearchServices searches services in MongoDB based on a keyword
// @Summary Search services
// @Description Search services
// @Produce json
// @Param keyword query string true "Search keyword"
// @Success 200 {array} Service
// @Router /services/search [get]
// @Tags Services
func SearchServices(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search keyword is required"})
		return
	}

	filter := bson.M{
		"$or": []bson.M{
			{"services_title": bson.M{"$regex": keyword, "$options": "i"}},
			{"services_body": bson.M{"$regex": keyword, "$options": "i"}},
			{"services_type": bson.M{"$regex": keyword, "$options": "i"}},
			{"services_assets": bson.M{"$regex": keyword, "$options": "i"}},
		},
	}

	var services []Service
	cursor, err := mongodb.DB.Collection("services").Find(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying services"})
		return
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &services); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning services"})
		return
	}

	c.JSON(http.StatusOK, services)
}
