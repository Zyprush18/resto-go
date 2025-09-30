package request

type User struct{
	Name	string 	`json:"name" validate:"required"`
	Email 	string	`json:"email" validate:"required,email"`
	Phone  	string 	`json:"phone" validate:"required"`
	Password string	`json:"password" validate:"required"`
	Role	string	`json:"role" validate:"required"`
}


type Menu struct{
	Name string `json:"name_menu" validate:"required"`
	Price int `json:"price" validate:"required"`
	Image string `json:"image" `
	IsAvailable *bool `json:"is_available"`	
}

type Order struct{
	TotalPrice int `json:"total_price"`
	Status  string `json:"status" validate:"required"`
	UserId uint `json:"user_id"`
}

type OrderItem struct{
	Quantity int `json:"quantity" validate:"required"`
	Price 	int `json:"price" validate:"required"`
	OrderId	uint `json:"order_id" validate:"required"`
	MenuId uint	`json:"menu_id" validate:"required"`
}

type Reservation struct{
	Date       string `json:"date_day" validate:"required"`
	Time       string `json:"time_day" validate:"required"`
	GuestCount int    `json:"guest_count" validate:"required"`
	Status     string `json:"status" validate:"required"`
	UserId     uint   `json:"user_id"`
}