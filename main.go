package main

import (
	// "net/http"
	// Importing features
	"net/http"

	// "github.com/gin-contrib/cors"
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

	router.Use(enableCORS())

	// Serve static files (HTML, CSS, JS, etc.)
	router.Static("/static", "./static")
	router.Static("/docs", "./docs")

	// // Define routes
	// router.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "Hello, World!")
	// })

	// Prepare a fallback route to always serve the 'index.html', had there not be any matching routes.

	// Serve static files from the "./web/build" directory.
	// router.StaticFS("/", http.Dir("./templates/Portfolio-master/build"))

	// // Serve index.html for all non-static routes.
	// router.GET("/", func(c *gin.Context) {
	// 	c.File("./templates/Portfolio-master/build/index.html")
	// })

	// Handle 404 errors
	router.NoRoute(func(c *gin.Context) {
		c.File("./templates/404.html")
	})

	// Define API versioning group
	v1 := router.Group("/api/v1")
	{
		v1.Use(enableCORS())

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
			data.Use(enableCORS())

			data.GET("/", alldata.GetAllData)
			data.GET("/init", alldata.PostAllData)
			data.DELETE("/", alldata.DeleteAllData)
			data.POST("/messages", alldata.SendMessage)

		}

	}

	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
	// defer sqldb.DB.Close()
}

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Content-Type", "application/json")
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://golang-my-portfolio-backend.onrender.com")
// 		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 		} else {
// 			c.Next()
// 		}
// 		c.Next()
// 	}
// }

// // Initialize the CORS middleware with a config object
// router.Use(cors.New(cors.Config{
// 	AllowOrigins: []string{"https://golang-my-portfolio-backend.onrender.com"},
// 	// Specifies the origins allowed to access the API
// 	AllowMethods: []string{"PUT", "PATCH"},
// 	// Specifies the allowed HTTP methods for cross-origin requests
// 	AllowHeaders: []string{"Origin"},
// 	// Specifies additional headers allowed for cross-origin requests
// 	ExposeHeaders: []string{"Content-Length"},
// 	// Specifies the headers that can be exposed from the response
// 	AllowCredentials: true,
// 	// Allows sending credentials (cookies, HTTP Authentication, and client-side SSL certificates)
// 	AllowOriginFunc: func(origin string) bool {
// 		// Allow cross-origin requests from the specified origin
// 		// Used for handling dynamic origin values or additional checks
// 		return origin == "https://github.com"
// 	},
// 	MaxAge: 12 * time.Hour,
// 	// Sets the inactivity timeout for pre-flight requests (default 12 hours)
// }))

// enableCORS is a middleware to enable CORS
func enableCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Allow requests from any origin
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// Allow specified HTTP methods
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		// Allow specified headers
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")

		// Check if it's a preflight request and handle it
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		// Continue with the next handler
		c.Next()
	}
}
