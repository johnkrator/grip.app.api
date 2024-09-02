package middleware

import (
	"github.com/gin-gonic/gin"
	"grip.app.api/api"
	config2 "grip.app.api/config"
	"grip.app.api/internal/controller"
	"grip.app.api/internal/models"
	"grip.app.api/internal/repository/account_repo"
	"grip.app.api/internal/repository/user_profile_repo"
	"grip.app.api/internal/repository/user_repo"
	"grip.app.api/internal/service/user_service"
)

func Run() error {
	db, err := config2.SetupDatabase()
	if err != nil {
		return err
	}

	models.SetupRelations(db)

	userRepo := user_repo.NewUserRepository(db)
	userProfileRepo := user_profile_repo.NewUserProfileRepository(db)
	accountRepo := account_repo.NewAccountRepository(db)

	newUserService := user_service.NewUserService(userRepo, userProfileRepo, accountRepo)
	userController := controller.NewUserController(newUserService)

	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	api.SetupRoutes(r, userController, newUserService)

	return r.Run(":8080")
}
