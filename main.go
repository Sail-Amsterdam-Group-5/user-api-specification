package main

import (
    "github.com/gin-gonic/gin"
    "user-api-specification/controllers"
	"user-api-specification/middleware"
	"user-api-specification/docs" // Swagger docs
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
)

// @title User Microservice API
// @version 1.0
// @description This is a User microservice with CRUD operations in GoLang.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.admin Grants read and write access to administrative information
// @scope.volunteer Volunteer access
// @scope.teamLead Team lead access
// @scope.keycloak Keycloak access


func main() {
    router := gin.Default()

	// Initialize Swagger doc info
    docs.SwaggerInfo.BasePath = "/"

    // Swagger endpoint
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // User CRUD routes
    user := router.Group("/user")
    {
        user.POST("/", middleware.CheckScope("admin"), controllers.CreateUser)
        user.GET("/:id", middleware.CheckScope("volunteer"), controllers.GetUser)
        user.PUT("/:id", middleware.CheckScope("team-lead"), controllers.UpdateUser)
        user.DELETE("/:id", middleware.CheckScope("admin"), controllers.DeleteUser)
    }

    router.Run(":8080")
}
