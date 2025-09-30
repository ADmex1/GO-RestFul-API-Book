package model

import (
	"time"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type Book struct {
	ID          uint      `json:"id" db:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" db:"book_title"`
	Slug        string    `json:"slug" db:"slug" gorm:"uniqueIndex"`
	Author      string    `json:"author" db:"book_author"`
	Description string    `json:"description" db:"book_description"`
	FileUpload  string    `json:"file_location" column:"file_location"`
	CreatedBy   int64     `json:"created_by" gorm:"not_null"`
	User        User      `json:"user" gorm:"foreignKey:CreatedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AddedDate   time.Time `json:"added_date" db:"added_date" gorm:"autoCreateTime"`
}

func (Book *Book) BeforeCreate(tx *gorm.DB) (err error) {
	if Book.Slug == "" {
		Book.Slug = slug.Make(Book.Title)
	}
	return
}
