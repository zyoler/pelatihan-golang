package auth

import (
	"fmt"
	"log"
	"net/http"
	"notification/controller"
	"notification/models"
	"reflect"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func ApiKey(key http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api := r.Header.Get("x-api-key")
		if api != "Golang" {
			controller.ResponseApi(w, 401, nil, "Invalid Api Key")
			return
		}
		key.ServeHTTP(w, r)
	})
}

func ExtractToken(String string) string {
	split := strings.Split(String, " ")
	if len(split) == 2 {
		return split[1]
	}
	return ""
}

func GetToken(token http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if reflect.ValueOf(auth).IsZero() {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokens := ExtractToken(auth)

		_, err := jwt.Parse(tokens, func(token *jwt.Token) (interface{}, error) {
			aut := "ITEMKU"
			checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aut, false)
			if !checkAudience {
				w.WriteHeader(http.StatusUnauthorized)
				return nil, fmt.Errorf("Invalid Token")
			}
			Iss := "ITEMKU.ISS"
			checkIssuer := token.Claims.(jwt.MapClaims).VerifyIssuer(Iss, false)
			if !checkIssuer {
				w.WriteHeader(http.StatusUnauthorized)
				return nil, fmt.Errorf("Invalid Issuer")
			}
			return []byte("LOGIN_SECRET"), nil
		})
		if err != nil {
			return
		}
		token.ServeHTTP(w, r)
	})
}

func GenerateTokens(w http.ResponseWriter, r *http.Request) {

	timesExpired := time.Now().Add(30 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":    "ITEMKU.ISS",
		"Audience":  "ITEMKU",
		"ExpiresAt": timesExpired.Unix(),
	})

	tokenString, err := token.SignedString([]byte("LOGIN_SECRET"))

	fmt.Println(tokenString, err)
	log.Println(tokenString)
	if err != nil {
		log.Println("Error", err)
		controller.ResponseApi(w, 500, nil, "Error Generate JWT !!")
		return
	}

	var tokensMaps models.Tokens
	tokensMaps.Tokens = tokenString
	controller.ResponseApi(w, 200, tokensMaps, "Sukses Generate JWT !!")
}
