package requests

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginUserRequest struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100)" json:"email" binding:"required"`
	Password string `gorm:"type:varchar(50)" json:"password" binding:"required"`
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Active   bool   `gorm:"default:false" json:"active"`
}

// ComparePassword -> Compare hashed password with given password
func (user *LoginUserRequest) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (LoginUserRequest) TableName() string {
	return "users"
}
