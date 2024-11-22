package entity

import (
	"time"

	"github.com/Zyprush18/resto/model/response"
	"gorm.io/gorm"
)

type User struct{
	ID      uint  `json:"id"`
	Name	string 	`json:"name" `
	Email 	string	`json:"email" `
	Phone  	string 	`json:"phone" `
	Password string	`json:"-" gorm:"column:password"`
	Order  	[]response.Order` gorm:"foreignKey:UserId;references:id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index, column:deletedAt"`
}


type Menu struct{
	ID uint `json:"id"`
	Name string `json:"name_menu"`
	Price int `json:"price"`
	Image string `json:"image"`
	IsAvailable *bool `json:"is_available"`	
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index, column:deletedAt"`
}

type Order struct{
	ID uint `json:"id"`
	TotalPrice int `json:"total_price"`
	Status  string `json:"status"`
	UserId uint `json:"user_id"`
	User 	response.User`json:"user"`	
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index, column:deletedAt"`
}