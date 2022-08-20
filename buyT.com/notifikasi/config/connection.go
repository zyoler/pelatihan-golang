package config

import (
	"log"
	"notifikasi/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error Load Config Files !! ")
	}
}

func Connect() *gorm.DB {
	url := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	connect := user + ":" + password + "@tcp(" + url + ":" + port + ")/" + dbname
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(models.Notif{})
	return db
}
