package response

import "time"

type User struct{
	ID      uint  `json:"id"`
	Name	string 	`json:"name" `
	Email 	string	`json:"email" `
	Phone  	string 	`json:"phone" `
	Password string	`json:"-" gorm:"column:password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Order struct{
	ID uint `json:"id"`
	TotalPrice int `json:"total_price"`
	Status  string `json:"status"`
	UserId uint `json:"user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Menu struct{
	ID uint `json:"id"`
	Name string `json:"name_menu"`
	Price int `json:"price"`
	IsAvailable string `json:"is_available"`
	Image string `json:"image"`	
	CreatedAt time.Time
	UpdatedAt time.Time

}

type OrderItem struct{
	ID uint `json:"id"`
	Quantity int `json:"quantity"`
	Price 	int `json:"price"`
	OrderId	uint `json:"order_id"`
	MenuId uint	`json:"menu_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Reservation struct{
	ID         uint   `gorm:"primaryKey"`
	Date       string `json:"date_day"`
	Time       string `json:"time_day"`
	GuestCount int    `json:"guest_count"`
	Status     string `json:"status"`
	UserId     uint   `json:"user_id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}