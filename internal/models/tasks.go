package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Tasks struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description sql.NullString
	Status      string `gorm:"not null;default:'pending'"`
	DueDate     *time.Time
	UserID      uint 
	CategoryID  sql.NullInt64
	CompletedAt sql.NullTime
	Reminder    *time.Time
	Notes       sql.NullString
	User        User `gorm:"foreignKey:UserID"`
}

