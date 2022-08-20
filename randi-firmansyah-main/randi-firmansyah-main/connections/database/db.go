package database

import (
	"log"
	"randi_firmansyah/models/productModel"
	"randi_firmansyah/models/userModel"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connected *gorm.DB

func init() {
	Connected = connectToDb()
}

func connectToDb() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/db_product_go_ojak"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&productModel.Product{}, &userModel.User{})
	log.Println("Database Connected")
	return db
}
