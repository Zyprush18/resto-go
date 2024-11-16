package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct{
	Name	string 	`json:"name" `
	Email 	string	`json:"email" `
	Phone  	string 	`json:"phone" `
	Password string	`json:"-" gorm:"column:password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index, column:deletedAt"`
}