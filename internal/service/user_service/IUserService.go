package user_service

import (
	"github.com/google/uuid"
	"grip.app.api/internal/dtos/request"
	"grip.app.api/internal/dtos/response"
)

type IUserService interface {
	CreateUser(req *request.UserRegistrationRequestDto) (*response.UserRegistrationResponseDto, error)
	LoginUser(req *request.UserLoginRequestDto) (*response.UserLoginResponseDto, error)
	VerifyEmail(email, token string) error
	ForgotPassword(email string) error
	ResetPassword(req *request.ResetPasswordRequestDto) error
	ChangePassword(userID uuid.UUID, req *request.ChangePasswordRequestDto) error
	GetCurrentUser(userID uuid.UUID) (*response.UserLoginResponseDto, error)
	DeleteUser(userID uuid.UUID) error
	GetUser(userID uuid.UUID) (*response.UserLoginResponseDto, error)
	GetAllUsers(page, pageSize int) ([]*response.UserLoginResponseDto, int64, error)
	UpdateUser(userID uuid.UUID, req *request.UpdateUserRequestDto) (*response.UserResponseDto, error)
}
