package requests

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginUserRequest struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100)" json:"email" binding:"required"`
	Password string `gorm:"type:varchar(50)" json:"-"`
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Active   bool   `gorm:"default:false" json:"active"`
}

func (user *LoginUserRequest) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
func (user *LoginUserRequest) VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
func (LoginUserRequest) TableName() string {
	return "users"
}
