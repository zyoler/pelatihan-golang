package main

import (
	"fmt"
	"log"
	"net/http"
	"pertemuan_6/connection"
	"pertemuan_6/controller"
	"pertemuan_6/repository"
	"pertemuan_6/usecase"
	"reflect"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error Load Config Files !! ")
	}
}
func main() {
	r := chi.NewRouter()

	koneksidb := connection.ConnectToDb()
	repo := repository.NewRepo(koneksidb)
	usecase := usecase.NewUsecase(repo)
	ctrl := controller.NewController(usecase)

	r.Group(func(rs chi.Router) {
		// rs.Use(AuthorizationApikey, GetToken)
		rs.Get("/getdata", ctrl.GetData)
		rs.Post("/postdata", ctrl.PostData)
		rs.Delete("/delete-data-mhs/{id}", ctrl.DeleteData)
		rs.Put("/update-data-mhs/{id}", ctrl.UpdateData)

		rs.Get("/getdata-detail/{id}", ctrl.GetDataMhsDetail)
	})
	r.Group(func(rst chi.Router) {
		rst.Use(AuthorizationApikey)
		rst.Post("/login", GenerateTokens)
	})
	fmt.Println("Running Service")

	if err := http.ListenAndServe(":5000", r); err != nil {
		fmt.Println("Error Starting Service")
	}
	fmt.Println("Starting Services")
}

func AuthorizationApikey(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apikey := r.Header.Get("x-api-key")
		if apikey != "BANDUNGLALA" {
			log.Println("401 API KEY INVALID  !!")
			controller.ResponseApi(w, 401, nil, "API KEY INVALID  !!")
			return
		}
		f.ServeHTTP(w, r)
	})
}

func GetToken(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("authorization")

		if reflect.ValueOf(authorization).IsZero() {
			// Handling pertama Jika Header Authorization nya kosong
			log.Println("unauthorized !! ")
		}

		tokensreal := ExtractTokens(authorization)

		_, err := jwt.Parse(tokensreal, func(token *jwt.Token) (interface{}, error) {
			aud := "PASIM"
			CheckAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
			if !CheckAudience {
				return nil, fmt.Errorf("Invalid Audience")
			}
			iss := "PASIM.ISS"
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

// func Verify

func ExtractTokens(token string) string {
	///
	/// Bearer asdhaskdkasdhsagdhasdgjasgdhasdghadgjadsaj
	strarr := strings.Split(token, " ")
	if len(strarr) == 2 {
		return strarr[1]
	}
	return ""
}

func GenerateTokens(w http.ResponseWriter, r *http.Request) {

	timesExpired := time.Now().Add(30 * time.Minute) // expired date
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":    "PASIM.ISS",
		"Audience":  "PASIM",
		"ExpiresAt": timesExpired.Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
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
