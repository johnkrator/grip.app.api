package user_service

import (
	"grip.app.api/internal/dtos/request"
	"grip.app.api/internal/dtos/response"
)

type IUserService interface {
	CreateUser(req *request.UserRegistrationRequestDto) (*response.UserRegistrationResponseDto, error)
	LoginUser(req *request.UserLoginRequestDto) (*response.UserLoginResponseDto, error)
	VerifyEmail(email, token string) error
}
