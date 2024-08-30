package models

import (
	"golang.org/x/crypto/bcrypt"
	"grip.app.api/internal/models/base"
	"time"
)

type User struct {
	base.BaseModel
	FirstName   string
	LastName    string
	Email       string `gorm:"uniqueIndex"`
	PhoneNumber string
	DateOfBirth time.Time
	Address     string
	Password    string
	Role        string
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
