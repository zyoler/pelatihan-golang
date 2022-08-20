package connection

import (
	"log"
	"penerepan_jwt/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDb() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/restorant"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.Product{}, &models.User{}, &models.Kategori{})
	// db.AutoMigrate(&models.Pesanan{})
	return db
}
