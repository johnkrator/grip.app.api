package user_service

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
	"grip.app.api/internal/dtos/request"
	"grip.app.api/internal/dtos/response"
	"grip.app.api/internal/middleware/email_send_config"
	"grip.app.api/internal/models"
	"grip.app.api/internal/repository/account_repo"
	"grip.app.api/internal/repository/user_profile_repo"
	"grip.app.api/internal/repository/user_repo"
	"grip.app.api/utils"
	"math/rand"
	"os"
	"time"
)

type UserService struct {
	userRepo        user_repo.IUserRepository
	userProfileRepo user_profile_repo.IUserProfileRepository
	accountRepo     account_repo.IAccountRepository
}

func NewUserService(userRepo user_repo.IUserRepository, userProfileRepo user_profile_repo.IUserProfileRepository, accountRepo account_repo.IAccountRepository) *UserService {
	return &UserService{
		userRepo:        userRepo,
		userProfileRepo: userProfileRepo,
		accountRepo:     accountRepo,
	}
}

func (s *UserService) CreateUser(req *request.UserRegistrationRequestDto) (*response.UserRegistrationResponseDto, error) {
	// Check if email already exists
	existingUser, _ := s.userRepo.GetUserByEmail(req.Email)
	if existingUser != nil {
		utils.ErrorLogger.Printf("Attempt to create user with existing email: %s", req.Email)
		return nil, utils.ErrEmailAlreadyExists
	}

	tx := s.userRepo.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user := &models.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		DateOfBirth: req.DateOfBirth,
		Address:     req.Address,
		Password:    req.Password,
		Role:        models.CustomerRole,
	}

	if err := user.HashPassword(); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := s.userRepo.CreateUserTx(tx, user); err != nil {
		tx.Rollback()
		return nil, err
	}

	userProfile := &models.UserProfile{
		UserID:        user.ID,
		Occupation:    req.Occupation,
		IncomeRange:   models.IncomeRange(req.IncomeRange),
		RiskTolerance: req.RiskTolerance,
		Preferences:   req.Preferences,
	}

	if err := s.userProfileRepo.CreateUserProfileTx(tx, userProfile); err != nil {
		tx.Rollback()
		return nil, err
	}

	account := &models.Account{
		UserID:        user.ID,
		AccountType:   models.SavingsAccount,
		AccountNumber: generateAccountNumber(),
		Balance:       decimal.NewFromFloat(0),
		Currency:      models.USD,
		Status:        models.Active,
		InterestRate:  decimal.NewFromFloat(0.01),
	}

	if err := s.accountRepo.CreateAccountTx(tx, account); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Generate 6-digit token
	token, err := s.generateToken(user.ID)
	if err != nil {
		utils.ErrorLogger.Printf("Failed to generate token: %v", err)
		return nil, err
	}

	// Send registration email with token and account number
	err = email_send_config.SendRegistrationEmail(user.Email, user.FirstName, user.LastName, account.AccountNumber, token)
	if err != nil {
		utils.ErrorLogger.Printf("Failed to send registration email: %v", err)
		// Log more details about the environment variables
		utils.ErrorLogger.Printf("SMTP_SERVER: %s", os.Getenv("SMTP_SERVER"))
		utils.ErrorLogger.Printf("SMTP_PORT: %s", os.Getenv("SMTP_PORT"))
		utils.ErrorLogger.Printf("SMTP_USER: %s", os.Getenv("SMTP_USER"))
		utils.ErrorLogger.Printf("FROM_EMAIL: %s", os.Getenv("FROM_EMAIL"))
	}

	return &response.UserRegistrationResponseDto{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		DateOfBirth: user.DateOfBirth,
		Address:     user.Address,
		Role:        response.Role(user.Role),
		Profile: response.UserProfileResponseDto{
			Occupation:    userProfile.Occupation,
			IncomeRange:   string(userProfile.IncomeRange),
			RiskTolerance: userProfile.RiskTolerance,
			Preferences:   userProfile.Preferences,
		},
		Account: response.AccountResponseDto{
			AccountNumber: account.AccountNumber,
			AccountType:   string(account.AccountType),
			Balance:       account.Balance,
			Currency:      string(account.Currency),
			Status:        string(account.Status),
		},
	}, nil
}

func (s *UserService) LoginUser(req *request.UserLoginRequestDto) (*response.UserLoginResponseDto, error) {
	user, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		utils.ErrorLogger.Printf("Failed to get user by email: %v", err)
		return nil, utils.ErrInvalidCredentials
	}

	if !user.IsVerified {
		utils.ErrorLogger.Printf("User not verified: %s", req.Email)
		return nil, errors.New("email not verified")
	}

	if err := user.ComparePassword(req.Password); err != nil {
		utils.ErrorLogger.Printf("Invalid password for user: %s", req.Email)
		return nil, utils.ErrInvalidCredentials
	}

	accessToken, refreshToken, err := generateTokens(user)
	if err != nil {
		utils.ErrorLogger.Printf("Failed to generate tokens: %v", err)
		return nil, err
	}

	// Update user with new tokens
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken

	// Save tokens to database
	err = s.userRepo.UpdateUser(user)
	if err != nil {
		utils.ErrorLogger.Printf("Failed to save tokens: %v", err)
		return nil, err
	}

	utils.InfoLogger.Printf("User logged in successfully: %s", user.Email)

	return &response.UserLoginResponseDto{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		DateOfBirth:  user.DateOfBirth,
		Address:      user.Address,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		IsVerified:   user.IsVerified,
		IsAdmin:      user.IsAdmin,
		IsDeleted:    user.IsDeleted,
		Role:         response.Role(user.Role),
	}, nil
}

func (s *UserService) VerifyEmail(email, token string) error {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return err
	}

	if user.Token != token {
		return errors.New("invalid token")
	}

	if time.Now().After(user.TokenExpiration) {
		return errors.New("token has expired")
	}

	user.IsVerified = true
	user.Token = ""
	user.TokenExpiration = time.Time{}

	err = s.userRepo.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) ChangePassword(userID uuid.UUID, req *request.ChangePasswordRequestDto) error {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return utils.ErrUserNotFound
	}

	// Verify current password
	if err := user.ComparePassword(req.CurrentPassword); err != nil {
		return utils.ErrInvalidCredentials
	}

	// Update password
	user.Password = req.NewPassword
	if err := user.HashPassword(); err != nil {
		return err
	}

	// Save updated user
	if err := s.userRepo.UpdateUser(user); err != nil {
		return err
	}

	return nil
}

func (s *UserService) ForgotPassword(email string) error {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return utils.ErrUserNotFound
	}

	// Generate a reset token
	resetToken, err := utils.GenerateRandomToken(32)
	if err != nil {
		return err
	}

	// Set token expiration (e.g., 1 hour from now)
	tokenExpiration := time.Now().Add(1 * time.Hour)

	// Update user with reset token and expiration
	user.ResetPasswordToken = resetToken
	user.ResetPasswordExpires = tokenExpiration

	if err := s.userRepo.UpdateUser(user); err != nil {
		return err
	}

	// Send password reset email
	resetLink := "https://localhost:8080/reset-password?token=" + resetToken
	err = email_send_config.SendPasswordResetEmail(user.Email, user.FirstName, resetLink)
	if err != nil {
		utils.ErrorLogger.Printf("Failed to send password reset email: %v", err)
		return err
	}

	return nil
}

func (s *UserService) ResetPassword(req *request.ResetPasswordRequestDto) error {
	user, err := s.userRepo.GetUserByResetToken(req.Token)
	if err != nil {
		return utils.ErrInvalidToken
	}

	if time.Now().After(user.ResetPasswordExpires) {
		return utils.ErrTokenExpired
	}

	// Update user's password
	user.Password = req.NewPassword
	if err := user.HashPassword(); err != nil {
		return err
	}

	// Clear reset token fields
	user.ResetPasswordToken = ""
	user.ResetPasswordExpires = time.Time{}

	if err := s.userRepo.UpdateUser(user); err != nil {
		return err
	}

	return nil
}

func generateAccountNumber() string {
	// Create a new random number generator with a seed based on the current time
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Generate a random number for the last 7 digits
	lastSevenDigits := r.Intn(10000000) // 7-digit number from 0000000 to 9999999

	// Combine the fixed "077" prefix with the random 7-digit number
	return fmt.Sprintf("001%07d", lastSevenDigits)
}

func init() {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

func generateTokens(user *models.User) (string, string, error) {
	// Access token claims
	accessClaims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 24 hour expiration
	}

	// Refresh token claims
	refreshClaims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 day expiration
	}

	// Create the access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
	if err != nil {
		return "", "", err
	}

	// Create the refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("REFRESH_TOKEN_SECRET")))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func (s *UserService) generateToken(userID uuid.UUID) (string, error) {
	token := fmt.Sprintf("%06d", rand.Intn(1000000))   // 6-digit random number
	expirationTime := time.Now().Add(15 * time.Minute) // 15 minutes expiration

	// Update user token with new token and expiration time in database table "users" and "user_tokens" tables
	// respectively using the user ID as the primary key and the token as the unique identifier for each token pair
	err := s.userRepo.UpdateUserToken(userID, token, expirationTime)
	if err != nil {
		return "", err
	}

	return token, nil
}
