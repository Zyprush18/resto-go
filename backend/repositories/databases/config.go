package databases

import (
	"fmt"
	"log"

	"github.com/Zyprush18/resto-go/backend/repositories/migration"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Migrate() {
	var err error
	DSN := "root:root@tcp(127.0.0.1:3306)/resto?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}
	var (
		User        migration.User
		Menu        migration.Menu
		Order       migration.Order
		OrderItem   migration.OrderItem
		Reservation migration.Reservation
	)

	errs := DB.AutoMigrate(&User, &Menu, &Order, &OrderItem, &Reservation)

	if errs != nil {
		log.Println(errs)
		fmt.Println("Gagal Migrate")
	}

	fmt.Println("Berhasil Connect")
	fmt.Println("Berhasil Migrate")
}
