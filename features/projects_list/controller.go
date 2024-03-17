package Projects

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

// @Summary Get all projects
// @Description Retrieve all projects from MongoDB
// @Tags Projects
// @Produce json
// @Success 200 {array} Project
// @Router /projects [get]
func GetProjects(c *gin.Context) {
	var projects []Project
	cursor, err := mongodb.DB.Collection("projects_list").Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error querying projects:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying projects"})
		return
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &projects); err != nil {
		log.Println("Error scanning projects:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning projects"})
		return
	}

	c.JSON(http.StatusOK, projects)
}

// @Summary Get a project by ID
// @Description Retrieve a project by its ID from MongoDB
// @Tags Projects
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {object} Project
// @Router /projects/{id} [get]
func GetProjectByID(c *gin.Context) {
	id := c.Param("id")
	var project Project
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	err = mongodb.DB.Collection("projects_list").FindOne(context.Background(), bson.M{"_id": objID}).Decode(&project)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	c.JSON(http.StatusOK, project)
}

// @Summary Create a new project
// @Description Create a new project in MongoDB
// @Tags Projects
// @Accept json
// @Produce json
// @Param project body Project true "Project object"
// @Success 201 {string} string "ID of the created project"
// @Router /projects [post]
func CreateProject(c *gin.Context) {
	var project Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := mongodb.DB.Collection("projects_list").InsertOne(context.Background(), project)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating project"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}

// @Summary Update a project
// @Description Update an existing project in MongoDB
// @Tags Projects
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Param project body Project true "Updated project object"
// @Success 200
// @Router /projects/{id} [put]
func UpdateProject(c *gin.Context) {
	id := c.Param("id")
	var project Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	_, err = mongodb.DB.Collection("projects_list").ReplaceOne(context.Background(), bson.M{"_id": objID}, project)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating project"})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Delete a project
// @Description Delete a project by its ID from MongoDB
// @Tags Projects
// @Param id path string true "Project ID"
// @Success 200
// @Router /projects/{id} [delete]
func DeleteProject(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	_, err = mongodb.DB.Collection("projects_list").DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting project"})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Delete all projects
// @Description Delete all projects from MongoDB
// @Tags Projects
// @Success 204
// @Router /projects [delete]
func DeleteAllProjects(c *gin.Context) {
	_, err := mongodb.DB.Collection("projects_list").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting projects"})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary Search projects
// @Description Search projects in MongoDB based on a keyword
// @Tags Projects
// @Produce json
// @Param keyword query string true "Search keyword"
// @Success 200 {array} Project
// @Router /projects/search [get]
func SearchProjects(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search keyword is required"})
		return
	}

	filter := bson.M{
		"$or": []bson.M{
			{"pl_title": bson.M{"$regex": keyword, "$options": "i"}},
			{"pl_body": bson.M{"$regex": keyword, "$options": "i"}},
		},
	}

	var projects []Project
	cursor, err := mongodb.DB.Collection("projects_list").Find(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying projects"})
		return
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &projects); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning projects"})
		return
	}

	c.JSON(http.StatusOK, projects)
}
