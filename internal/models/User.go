package models

import (
	"golang.org/x/crypto/bcrypt"
	"grip.app.api/internal/models/base"
	"time"
)

type Role string

const (
	AdminRole    Role = "admin"
	ManagerRole  Role = "manager"
	EmployeeRole Role = "employee"
	CustomerRole Role = "customer"
)

type User struct {
	base.BaseModel
	FirstName       string
	LastName        string
	Email           string `gorm:"uniqueIndex"`
	PhoneNumber     string
	DateOfBirth     time.Time
	Address         string
	Password        string
	AccessToken     string `gorm:"default:null"`
	RefreshToken    string `gorm:"default:null"`
	IsVerified      bool   `gorm:"default:false"`
	IsAdmin         bool   `gorm:"default:false"`
	IsDeleted       bool   `gorm:"default:false"`
	Role            Role   `gorm:"default:customer"`
	Token           string `gorm:"default:null"`
	TokenExpiration time.Time
}

func (baseUser *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(baseUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	baseUser.Password = string(hashedPassword)
	return nil
}

func (baseUser *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(baseUser.Password), []byte(password))
}
