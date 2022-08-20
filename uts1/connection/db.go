package connection

import (
	"log"
	"uts_1/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	// dbc = database connection
	dbc := "root:@tcp(127.0.0.1:3306)/db_projectgo"
	db, err := gorm.Open(mysql.Open(dbc), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.HeroML{}, &models.Books{})
	return db
}
