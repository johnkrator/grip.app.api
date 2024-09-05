package utils

import (
	"grip.app.api/internal/dtos/response"
	"grip.app.api/internal/models"
)

func MapUserToResponseDto(user *models.User) *response.UserLoginResponseDto {
	return &response.UserLoginResponseDto{
		ID:           user.ID,
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Role:         response.Role(user.Role),
		IsVerified:   user.IsVerified,
		IsAdmin:      user.IsAdmin,
		IsDeleted:    user.IsDeleted,
		PhoneNumber:  user.PhoneNumber,
		DateOfBirth:  user.DateOfBirth,
		Address:      user.Address,
		AccessToken:  user.AccessToken,
		RefreshToken: user.RefreshToken,
	}
}
