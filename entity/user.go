package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       int    		`json:"id"`
	Name     string 		`json:"name"`
	CreateAt time.Time 		`json:"created_at"`
	UpdateAt time.Time 		`json:"update_at"`
	DeleteAt gorm.DeletedAt `json:"_"`
}

func NewUser(name string) *User  {
	return &User{
		Name: name,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
}

func UpdateUser(id int, name string) *User  {
	return &User{
		ID: id,
		Name: name,
		UpdateAt: time.Now(),
	}
}