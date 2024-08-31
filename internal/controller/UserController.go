package controller

import (
	"grip.app.api/internal/dtos/request"
	"grip.app.api/internal/service/user_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService user_service.IUserService
}

func NewUserController(userService user_service.IUserService) *UserController {
	return &UserController{userService: userService}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param user body request.UserRegistrationRequestDto true "Register user"
// @Success 201 {object} response.UserRegistrationResponseDto
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /register [post]
func (c *UserController) Register(ctx *gin.Context) {
	var req request.UserRegistrationRequestDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.userService.CreateUser(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

// Login godoc
// @Summary Login a user
// @Description Login a user with the provided credentials
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body request.UserLoginRequestDto true "Login credentials"
// @Success 200 {object} response.UserLoginResponseDto
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Router /login [post]
func (c *UserController) Login(ctx *gin.Context) {
	var req request.UserLoginRequestDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.userService.LoginUser(&req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// Welcome godoc
// @Summary Welcome message
// @Description Get a welcome message
// @Tags welcome
// @Produce json
// @Success 200 {object} map[string]string
// @Router / [get]
func (c *UserController) Welcome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to the Grip API"})
}
