package model

import (
	"time"

	"gorm.io/gorm"
)

// gorm.Model definition
type Cat struct {
	ID uint `gorm:"primaryKey"`
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
  }