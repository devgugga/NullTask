package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string  `gorm:"not null"`
	Email        *string `gorm:"uniqueIndex"`
	Age          uint8   `gorm:"check:age >= 0"`
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Tasks        []Tasks
}
