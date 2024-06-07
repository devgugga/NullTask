package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           string     `json:"id" gorm:"type:uuid;primary_key;"`
	Name         string     `json:"name" validate:"required" gorm:"not null"`
	Email        string     `json:"email" validate:"required,email" gorm:"uniqueIndex"`
	Password     string     `json:"password" validate:"required,min=6" gorm:"not null"`
	Age          uint8      `json:"age" validate:"required,gte=18" gorm:"check:age >= 0"`
	MemberNumber int        `json:"member_number" gorm:"autoIncrement"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at" sql:"index"`
	Tasks        []Tasks    // `gorm:"foreignKey:UserId"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
