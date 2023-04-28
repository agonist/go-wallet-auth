package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Address string `gorm:"uniqueIndex" json:"address"`
}
