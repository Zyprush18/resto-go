package migration

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type User struct {
	ID          uint          `gorm:"primaryKey"`
	Name        string        `json:"name" gorm:"type:varchar(255)"`
	Email       string        `json:"email" gorm:"type:varchar(100);uniqueIndex;email"`
	Phone       string        `json:"phone," gorm:"type:varchar(12);uniqueIndex"`
	Password    string        `json:"password"`
	Order       []Order       `gorm:"foreignKey:UserId;references:id"`
	Reservation []Reservation `gorm:"foreignKey:UserId;references:id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Menu struct {
	ID          uint        `gorm:"primaryKey"`
	Name        string      `json:"name_menu"`
	Price       int         `json:"price"`
	IsAvailable string      `json:"is_available" gorm:"type:boolean"`
	Image       string      `json:"image"`
	OrderItem   []OrderItem `gorm:"foreignKey:MenuId;references:id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Order struct {
	ID         uint        `gorm:"primaryKey"`
	TotalPrice int         `json:"total_price"`
	Status     string      `json:"status" gorm:"column:status;type:enum('pending','processing','complete')"`
	UserId     uint        `json:"user_id"`
	OrderItem  []OrderItem `gorm:"foreignKey:OrderId;references:id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type OrderItem struct {
	ID        uint `gorm:"primaryKey"`
	Quantity  int  `json:"quantity"`
	Price     int  `json:"price"`
	OrderId   uint `json:"order_id"`
	MenuId    uint `json:"menu_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Reservation struct {
	ID         uint      `gorm:"primaryKey"`
	Date       string    `json:"date_day" gorm:"type:date"`
	Time       CustomTime `json:"time_day"`
	GuestCount int       `json:"guest_count"`
	Status     string    `json:"status" gorm:"column:status;type:enum('pending','processing','complete')"`
	UserId     uint      `json:"user_id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}


type CustomTime struct {
	time.Time
}

func (CustomTime) GormDataType() string {
    return "time"
}

func (CustomTime) GormDBDataType(db *gorm.DB, field *schema.Field) string {
    return "time"
}


