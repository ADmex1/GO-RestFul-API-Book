package model

import "time"

type Book struct {
	ID          uint      `json:"id" db:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" db:"book_title"`
	Author      string    `json:"author" db:"book_author"`
	Description string    `json:"description" db:"book_description"`
	AddedDate   time.Time `json:"added_date" db:"added_date" gorm:"autoCreateTime"`
}
