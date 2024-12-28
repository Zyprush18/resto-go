package request

type User struct{
	Name	string 	`json:"name" validate:"required"`
	Email 	string	`json:"email" validate:"required,email"`
	Phone  	string 	`json:"phone" validate:"required"`
	Password string	`json:"password" validate:"required"`
}


type Menu struct{
	Name string `json:"name_menu" validate:"required"`
	Price int `json:"price" validate:"required"`
	Image string `json:"image" `
	IsAvailable *bool `json:"is_available"`	
}

type Order struct{
	TotalPrice int `json:"total_price"`
	Status  string `json:"status"`
	UserId uint `json:"user_id"`
}

type OrderItem struct{
	Quantity int `json:"quantity"`
	Price 	int `json:"price"`
	OrderId	uint `json:"order_id"`
	MenuId uint	`json:"menu_id"`
}