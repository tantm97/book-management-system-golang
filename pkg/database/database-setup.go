package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/tantm97/book-management-system-golang/pkg/models"
)

var DB *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "test:password@/test_golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DB = d
	fmt.Println("connected database!")
	DB.AutoMigrate(&models.Book{}, &models.User{})
}
