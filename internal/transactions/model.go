package transactions

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	Name     string
	Email    string
	Code     string
	Password string
	Active   bool
	Roles    []Role `gorm:"many2many:user_role;"`
}

func (u *User) toResponse() *UserResponse {
	return &UserResponse{
		Code:  u.Code,
		Name:  u.Name,
		Email: u.Email,
	}
}

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserRequest) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *UserRequest) toUser() *User {
	return &User{
		Name:     u.Name,
		Email:    u.Email,
		Active:   false,
		Code:     "ET-" + uuid.NewString(),
		Password: u.Password,
	}
}

type UserResponse struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
