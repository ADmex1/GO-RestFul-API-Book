package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	InternalID int64          `json:"internal_id" gorm:"primaryKey"`
	PublicID   uuid.UUID      `json:"public_id" gorm:"type:uuid"`
	Name       string         `json:"name" db:"name"`
	Email      string         `json:"email" db:"email" gorm:"unique"`
	Password   string         `json:"-"  gorm:"column:password"`
	Role       string         `json:"role" gorm:"default:'user'"`
	CreatedAt  time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}
