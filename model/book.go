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
	AddedDate   time.Time `json:"added_date" db:"added_date" gorm:"autoCreateTime"`
}

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	if b.Slug == "" {
		b.Slug = slug.Make(b.Title)
	}
	return
}
