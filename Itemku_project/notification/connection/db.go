package connection

import (
	"log"
	"notification/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error Load Config Files !!")
	}
}

func ConnectToDb() *gorm.DB {
	url := os.Getenv("URL_DB")
	port := os.Getenv("PORT_DB")
	DB := os.Getenv("DB_NAME")
	user := os.Getenv("USERNAME_DB")
	pw := os.Getenv("PW_DB")

	log.Println("testin url", url)
	dsn := user + ":" + pw + "@tcp(" + url + ":" + port + ")/" + DB + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.NotifItemku{})
	return db
}
