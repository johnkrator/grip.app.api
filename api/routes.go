package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"grip.app.api/internal/controller"
	"grip.app.api/internal/handlers"
	"grip.app.api/internal/service/user_service"
	"grip.app.api/utils"
)

func SetupRoutes(r *gin.Engine, userController *controller.UserController, userService user_service.IUserService) {
	// Initialize UserHandler
	userHandler := handlers.NewUserHandler(userService)

	// Swagger route should be before other routes and error handlers
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Apply ErrorHandler middleware to the group of routes
	api := r.Group("/")
	api.Use(utils.ErrorHandler())
	{
		api.POST("/register", userController.Register)
		api.POST("/login", userController.Login)
		api.POST("/verify-email", userHandler.VerifyEmail)
	}
}
