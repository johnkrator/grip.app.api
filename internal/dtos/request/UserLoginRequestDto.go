package request

type UserLoginRequestDto struct {
	Email    string `gorm:"uniqueIndex"`
	Password string
}
