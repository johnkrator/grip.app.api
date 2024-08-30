package middleware

import (
	"github.com/gin-gonic/gin"
	config2 "grip.app.api/config"
	"grip.app.api/internal/models"
)

func Run() error {
	db, err := config2.SetupDatabase()
	if err != nil {
		return err
	}

	models.SetupRelations(db)

	//userRepo := userRepository.NewUserRepository(db)
	//newUserService := userService.NewUserService(userRepo)
	//userController := controllers.NewUserController(newUserService)

	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	//routes.SetupRoutes(r, userController, orderController)

	return r.Run(":8080")
}
