package auth

import (
	"net/http"

	"gorm.io/gorm"
)

type Tokens struct {
	Token string `json:"token"`
}

type BD struct {
	DB *gorm.DB
}

type authorization interface {
	GetToken(token http.Handler) http.Handler
	ExtractToken(r *http.Request) string
	ApiKey(w http.Handler) http.Handler
}
