package main

import (
	"net/http"
	// Importing features

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// sqldb "main/core/db/sql"
	mongodb "main/core/db/monodb"
	_ "main/docs" // This is required for Swagger to find your documentation
	services "main/features/Services"
	alldata "main/features/all_data"
	home_details "main/features/home_details"
	projects "main/features/projects_list"
	social_media "main/features/social_media"
)

// @title User API
// @description API for user management
// @version 1.0
// @host golang-my-portfolio-backend.onrender.com
// @BasePath /api/v1
func main() {

	// localhost:8080
	// golang-my-portfolio-backend.onrender.com
	// sqldb.Init()
	mongodb.InitMongoDB()
	// InitDataBase()

	router := gin.Default()

	// Serve static files (HTML, CSS, JS, etc.)
	router.Static("/static", "./static")
	router.Static("/docs", "./docs")

	// Define routes
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// Handle 404 errors
	router.NoRoute(func(c *gin.Context) {
		c.File("./templates/404.html")
	})

	// Define API versioning group
	v1 := router.Group("/api/v1")
	{
		// Grouping routes related to home details management
		hdRoute := v1.Group("/home_details")
		{
			// HomeDetails routes
			hdRoute.POST("/", home_details.CreateHomeDetails)
			hdRoute.PUT("/:id", home_details.UpdateHomeDetails)
			hdRoute.GET("/:id", home_details.GetHomeDetailsByID)
			hdRoute.GET("", home_details.GetHomeDetails)
			hdRoute.GET("/search", home_details.SearchHomeDetailses)
			hdRoute.DELETE("/:id", home_details.DeleteHomeDetails)
			hdRoute.DELETE("/", home_details.DeleteAllHomeDetailses)
		}

		// Grouping routes related to social media management
		smRoute := v1.Group("/social_media")
		{
			// SocialMedia routes
			smRoute.POST("/", social_media.CreateSocialMediaDetails)
			smRoute.PUT("/:id", social_media.UpdateSocialMediaDetails)
			smRoute.GET("/:id", social_media.GetSocialMediaDetailsByID)
			smRoute.GET("", social_media.GetSocialMediaDetails)
			smRoute.GET("/search", social_media.SearchSocialMediaDetails)
			smRoute.DELETE("/:id", social_media.DeleteSocialMediaDetails)
			smRoute.DELETE("/", social_media.DeleteAllSocialMediaDetails)
		}

		// Grouping routes related to projects management
		projectsRoute := v1.Group("/projects")
		{
			// Projects routes
			projectsRoute.POST("/", projects.CreateProject)
			projectsRoute.PUT("/:id", projects.UpdateProject)
			projectsRoute.GET("/:id", projects.GetProjectByID)
			projectsRoute.GET("", projects.GetProjects)
			projectsRoute.GET("/search", projects.SearchProjects)
			projectsRoute.DELETE("/:id", projects.DeleteProject)
			projectsRoute.DELETE("/", projects.DeleteAllProjects)
		}

		// Grouping routes related to services management
		sRoute := v1.Group("/services")
		{
			// Services routes
			sRoute.POST("/", services.CreateServices)
			sRoute.PUT("/:id", services.UpdateServices)
			sRoute.GET("/:id", services.GetServicesByID)
			sRoute.GET("", services.GetServices)
			sRoute.GET("/search", services.SearchServices)
			sRoute.DELETE("/:id", services.DeleteServices)
			sRoute.DELETE("/", services.DeleteAllServices)
		}
		data := v1.Group("/data")
		{
			data.GET("/", alldata.GetAllData)
			data.DELETE("/", alldata.DeleteAllData)
		}

	}

	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
	// defer sqldb.DB.Close()
}
