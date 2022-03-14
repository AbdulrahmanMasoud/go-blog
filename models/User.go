package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Email    string `gorm:"type:varchar(100)" json:"email"`
	Password []byte `gorm:"type:varchar(500)" json:"password"`
	Active   bool   `gorm:"default:false" json:"active"`
	Blogs    []Blog `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
