package transactions

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID   uint64 `gorm:"autoIncrement"`
	Role string `gorm:"unique"`
}

type User struct {
	gorm.Model
	ID       uint64 `gorm:"autoIncrement"`
	Code     string
	Name     string
	Email    string
	Password string
}
