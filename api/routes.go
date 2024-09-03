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

const (
	apiV1 = "/api/v1"
)

func SetupRoutes(r *gin.Engine, userController *controller.UserController, userService *user_service.UserService) {
	// Initialize UserHandler
	verifyUserEmailHandler := handlers.NewUserHandler(userService)

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API v1 routes
	v1 := r.Group(apiV1)
	v1.Use(utils.ErrorHandler())
	{
		// Public routes
		v1.POST("/register", userController.Register)
		v1.POST("/login", userController.Login)
		v1.POST("/verify-email", verifyUserEmailHandler.VerifyEmail)
		v1.POST("/forgot-password", userController.ForgotPassword)
		v1.POST("/reset-password", userController.ResetPassword)

		// Protected routes
		authorized := v1.Group("/")
		authorized.Use(handlers.AuthMiddleware())
		{
			authorized.POST("/change-password", userController.ChangePassword)

			// New protected routes
			authorized.GET("/users/me", userController.GetCurrentUser)
			authorized.GET("/users/:id", userController.GetUser)
			authorized.DELETE("/users/:id", userController.DeleteUser)
			authorized.GET("/users", userController.GetAllUsers)
			// ... other protected routes ...
		}
	}
}
