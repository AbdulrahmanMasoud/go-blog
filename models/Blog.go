package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	UserId  uint   `json:"user_id"`
	Title   string `gorm:"type:varchar(500)" json:"title"`
	Content string `gorm:"type:text" json:"content"`
	Slug    string `gorm:"type:varchar(500);unique" json:"slug"`
	Visible bool   `gorm:"default:true" json:"visible"`
}
