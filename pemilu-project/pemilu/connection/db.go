package connection

import (
	"pemilu/models"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	// Untuk sql
	// dsn := "root:@tcp(127.0.0.1:3306/db_belajar?"

	// dsn := "host=localhost user=postgres password=123456789 dbname=belajar port=5432 sslmode=disable"
	dsn := "postgres://postgres:123456789@localhost:5432/belajar"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Mahasiswa{})
	return db
}
