package requests

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterUserRequest struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100)" json:"name" binding:"required"`
	Email    string `gorm:"type:varchar(100)" json:"email" binding:"required"`
	Password string `gorm:"type:varchar(50)" json:"password" binding:"required"`
	Active   bool   `gorm:"default:false" json:"active"`
}

func (RegisterUserRequest) TableName() string {
	return "users"
}

func (user *RegisterUserRequest) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	user.Password = string(hashedPassword)
	return
}
