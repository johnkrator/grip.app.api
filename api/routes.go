package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"grip.app.api/internal/controller"
	"grip.app.api/utils"
)

func SetupRoutes(r *gin.Engine, userController *controller.UserController) {
	// Swagger route should be before other routes and error handlers
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Apply ErrorHandler middleware to the group of routes
	api := r.Group("/")
	api.Use(utils.ErrorHandler())
	{
		api.POST("/register", userController.Register)
		api.POST("/login", userController.Login)
	}
}
