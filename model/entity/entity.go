package entity

import (
	"time"

	"github.com/Zyprush18/resto-go/model/response"
	"gorm.io/gorm"
)

type User struct {
	ID          uint                   `json:"id"`
	Name        string                 `json:"name" `
	Email       string                 `json:"email" `
	Phone       string                 `json:"phone" `
	Password    string                 `json:"-" gorm:"column:password"`
	Order       []response.Order       ` gorm:"foreignKey:UserId;references:id"`
	Reservation []response.Reservation `gorm:"foreignKey:UserId;references:id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index, column:deletedAt"`
}

type Menu struct {
	ID          uint                 `json:"id"`
	Name        string               `json:"name_menu"`
	Price       int                  `json:"price"`
	Image       string               `json:"image"`
	IsAvailable *bool                `json:"is_available"`
	OrderItem   []response.OrderItem `gorm:"foreignKey:MenuId;references:id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index, column:deletedAt"`
}

type Order struct {
	ID         uint                 `json:"id"`
	TotalPrice int                  `json:"total_price"`
	Status     string               `json:"status"`
	UserId     uint                 `json:"user_id"`
	User       response.User        `json:"user"`
	OrderItem  []response.OrderItem ` gorm:"foreignKey:OrderId;references:id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index, column:deletedAt"`
}

type OrderItem struct {
	ID        uint           `json:"id"`
	Quantity  int            `json:"quantity"`
	Price     int            `json:"price"`
	OrderId   uint           `json:"order_id"`
	MenuId    uint           `json:"menu_id"`
	Order     response.Order `json:"order"`
	Menu      response.Menu  `json:"menu"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index, column:deletedAt"`
}

type Reservation struct {
	ID         uint          `json:"id"`
	Date       string        `json:"date_day"`
	Time       string        `json:"time_day"`
	GuestCount int           `json:"guest_count"`
	Status     string        `json:"status"`
	UserId     uint          `json:"user_id"`
	User       response.User `json:"user"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index, column:deleteAt"`
}
