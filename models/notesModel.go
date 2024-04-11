package models

import "gorm.io/gorm"

type Notes struct {
	gorm.Model
	Title       string
	IsCompleted bool
}
