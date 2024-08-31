package user_service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
	"grip.app.api/internal/dtos/request"
	"grip.app.api/internal/dtos/response"
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

func NewUserService(
	userRepo user_repo.IUserRepository,
	userProfileRepo user_profile_repo.IUserProfileRepository,
	accountRepo account_repo.IAccountRepository,
) *UserService {
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
		AccountType:   models.CurrentAccount,
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

func generateAccountNumber() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%010d", rand.Intn(9000000000)+1000000000)
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
