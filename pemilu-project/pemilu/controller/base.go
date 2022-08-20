package controller

import (
	"net/http"
	"pemilu/connection"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var DB *gorm.DB
var REDIS *redis.Client

func init() {
	DB = connection.ConnectToDB()
	REDIS = connection.ConnectToRedis()
}

type ControllerInterface interface {
	GetDataPenyelenggara(w http.ResponseWriter, r *http.Request)
}
