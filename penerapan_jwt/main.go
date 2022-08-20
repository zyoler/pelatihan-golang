package main

import (
	"fmt"
	"log"
	"net/http"
	"penerepan_jwt/controller"
	"reflect"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt"
)

func main() {
	Router := chi.NewRouter()
	Router.Group(func(r chi.Router) {
		r.Use(GetToken)
		r.Get("/getkatagori", controller.GetKatagoris)
		r.Post("/postkategori", controller.PostKategoris)
		r.Delete("/deletekategori/{id}", controller.DeleteKategoris)
		r.Put("/updatekategori/{id}", controller.UpdateKategoris)

		r.Get("/getproduct", controller.GetProduct)
		r.Post("/postproduct", controller.PostProduct)
		r.Delete("/deleteproduct/{id}", controller.DeleteProduct)
		r.Put("/updateproduct/{id}", controller.UpdateProduct)

		r.Get("/getuser", controller.GetUser)
		r.Post("/postuser", controller.PostUser)
		r.Delete("/deleteuser/{id}", controller.DeleteUser)
		r.Put("/updateuser/{id}", controller.UpdateUser)
	})

	Router.Group(func(r chi.Router) {
		r.Post("/login", GenerateTokens)
	})

	fmt.Println("Running Service")

	if err := http.ListenAndServe(":5000", Router); err != nil {
		fmt.Println("Error Starting Service")
	}
	fmt.Println("Starting Services")
}

func GetToken(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("authorization")

		if reflect.ValueOf(authorization).IsZero() {
			log.Println("unauthorized !! ")
		}

		tokensreal := ExtractTokens(authorization)

		_, err := jwt.Parse(tokensreal, func(token *jwt.Token) (interface{}, error) {
			aud := "JWT"
			CheckAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
			if !CheckAudience {
				return nil, fmt.Errorf("Invalid Audience")
			}
			iss := "JWT.ISS"
			CheckISS := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
			if !CheckISS {
				return nil, fmt.Errorf("Invalid ISS")
			}

			return []byte("LOGIN_SECRET"), nil
		})

		if err != nil {
			log.Println("401 unatuhorized !!")
			controller.ResponseApi(w, 401, nil, "Unauthorized !!")
			return
		}
		f.ServeHTTP(w, r)
	})

}
func ExtractTokens(token string) string {
	strarr := strings.Split(token, " ")
	if len(strarr) == 2 {
		return strarr[1]
	}
	return ""
}

func GenerateTokens(w http.ResponseWriter, r *http.Request) {

	timesExpired := time.Now().Add(15 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":    "PUB",
		"Audience":  "PASIM",
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

	type Tokens struct {
		Tokens string `json:'tokens'`
	}
	var tokensMaps Tokens
	tokensMaps.Tokens = tokenString
	controller.ResponseApi(w, 200, tokensMaps, "Sukses Generate JWT !!")

}
