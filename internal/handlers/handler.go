package handlers

import (
	"github.com/charmbracelet/log"
	"gorm.io/gorm"
)

type Handler struct {
	DB  *gorm.DB
	Log *log.Logger
}
