package controller

import (
	"github.com/google/uuid"
	"grip.app.api/internal/dtos/request"
	"grip.app.api/internal/dtos/response"
	"grip.app.api/internal/service/user_service"
	"grip.app.api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService user_service.IUserService
}

func NewUserController(userService *user_service.UserService) *UserController {
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

// ChangePassword godoc
// @Summary Change user's password
// @Description Change the authenticated user's password
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param changePasswordInfo body request.ChangePasswordRequestDto true "Change password information"
// @Success 200 {object} response.ChangePasswordResponseDto
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /change-password [post]
func (c *UserController) ChangePassword(ctx *gin.Context) {
	var req request.ChangePasswordRequestDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	err := c.userService.ChangePassword(userID.(uuid.UUID), &req)
	if err != nil {
		if err == utils.ErrInvalidCredentials {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid current password"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, response.ChangePasswordResponseDto{Message: "Password changed successfully"})
}

// ForgotPassword godoc
// @Summary Initiate password reset
// @Description Send a password reset email to the user
// @Tags users
// @Accept json
// @Produce json
// @Param email body request.ForgotPasswordRequestDto true "User's email"
// @Success 200 {object} response.ForgotPasswordResponseDto
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /forgot-password [post]
func (c *UserController) ForgotPassword(ctx *gin.Context) {
	var req request.ForgotPasswordRequestDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.userService.ForgotPassword(req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.ForgotPasswordResponseDto{Message: "Password reset email sent successfully"})
}

// ResetPassword godoc
// @Summary Reset user's password
// @Description Reset the user's password using the provided token
// @Tags users
// @Accept json
// @Produce json
// @Param resetInfo body request.ResetPasswordRequestDto true "Reset password information"
// @Success 200 {object} response.ResetPasswordResponseDto
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /reset-password [post]
func (c *UserController) ResetPassword(ctx *gin.Context) {
	var req request.ResetPasswordRequestDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.userService.ResetPassword(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.ResetPasswordResponseDto{Message: "Password reset successfully"})
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

// GetCurrentUser godoc
// @Summary Get current user
// @Description Get the current user
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} response.UserLoginResponseDto "User"
// @Failure 401 {object} gin.H
// @Router /current-user [get]
func (c *UserController) GetCurrentUser(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	uid, ok := userID.(uuid.UUID)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	user, err := c.userService.GetCurrentUser(uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete the user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} gin.H "User deleted successfully"
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/{id} [delete]
func (c *UserController) DeleteUser(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = c.userService.DeleteUser(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// GetUser godoc
// @Summary Get user
// @Description Get a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 2 {object} response.UserLoginResponseDto "User"
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/{id} [get]
func (c *UserController) GetUser(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := c.userService.GetUser(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Success 200 {object} response.UserLoginResponseDto "Users"
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users [get]
func (c *UserController) GetAllUsers(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))

	users, totalCount, err := c.userService.GetAllUsers(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users":       users,
		"total_count": totalCount,
		"page":        page,
		"page_size":   pageSize,
	})
}

// UpdateUser godoc
// @Summary Update user information
// @Description Update the authenticated user's information
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param updateInfo body request.UpdateUserRequestDto true "Update user information"
// @Success 200 {object} response.UserResponseDto
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/{id} [put]
func (c *UserController) UpdateUser(ctx *gin.Context) {
	var req request.UpdateUserRequestDto
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	updatedUser, err := c.userService.UpdateUser(userID, &req)
	if err != nil {
		switch err.Error() {
		case "user not found":
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		case "failed to update user", "failed to get user profile", "failed to update user profile":
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
		}
		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}
