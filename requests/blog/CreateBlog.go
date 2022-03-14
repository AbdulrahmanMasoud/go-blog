package requests

import (
	"gorm.io/gorm"
	"strings"
)

type CreateBlogRequest struct {
	gorm.Model
	UserId  uint   `json:"user_id" `
	Title   string `gorm:"type:varchar(500)" json:"title" binding:"required"`
	Content string `gorm:"type:text" json:"content" binding:"required"`
	Slug    string `gorm:"type:varchar(500);unique" json:"slug"`
}

func (CreateBlogRequest) TableName() string {
	return "blogs"
}

func (blog *CreateBlogRequest) BeforeCreate(tx *gorm.DB) (err error) {
	slug := strings.ToLower(strings.Replace(blog.Title, " ", "-", -1))
	blog.Slug = slug
	return
}
