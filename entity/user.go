package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"_"`
	CreatedAt time.Time      `json:"create_at"`
	UpdatedAt time.Time      `json:"_"`
	DeletedAt gorm.DeletedAt `json:"_"`
}

func NewUser(name, email, password string) *User {
	return &User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func UpdateUser(id int64, name, email, password string) *User {
	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  password,
		UpdatedAt: time.Now(),
	}
}
