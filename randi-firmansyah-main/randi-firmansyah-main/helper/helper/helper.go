package helper

import (
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
)

var (
	REDIS *redis.Client
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func UnMarshall(from string, to interface{}) error {
	if err := json.Unmarshal([]byte(from), &to); err != nil {
		return err
	}
	return nil
}
