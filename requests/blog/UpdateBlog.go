package requests

import (
	"gorm.io/gorm"
	"strings"
)

type UpdateBlogRequest struct {
	gorm.Model
	UserId  uint   `json:"user_id"`
	Title   string `gorm:"type:varchar(500)" json:"title" binding:"required"`
	Content string `gorm:"type:text" json:"content" binding:"required"`
	Slug    string `gorm:"type:varchar(500);unique" json:"slug"`
}

func (UpdateBlogRequest) TableName() string {
	return "blogs"
}

func (blog *UpdateBlogRequest) BeforeUpdate(tx *gorm.DB) (err error) {
	slug := strings.ToLower(strings.Replace(blog.Title, " ", "-", -1))
	blog.Slug = slug
	return
}
