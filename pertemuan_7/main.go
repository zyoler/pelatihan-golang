package main

import (
	"fmt"
	"log"
	"net/http"
	"pertmuan_2/controller"
	"reflect"
	"strings"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/getdata", controller.GetData)
	r.Post("/postdata", controller.PostData)
	r.Delete("/delete-data-mhs/{id}", controller.DeleteData)
	r.Put("/update-data-mhs/{id}", controller.UpdateData)

	fmt.Println("Running Service")

	if err := http.ListenAndServe(":5000", r); err != nil {
		fmt.Println("Error Starting Service")
	}
	fmt.Println("Starting Services")
}

func GetToken(f http.Handler)http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		authorization := r.Header.Get("authorization")

		if reflect.Valueof(authorization).isZero() {
			log.Println("unauthorized")
		}

		tokensreal := ExtractTokens(authorization)
		_, err := jwt.Parse(tokensreal, func(token *jwt.Token))
	})
}

func ExtractTokens(token string) string{
	strarr := strings.Split(token, " ")
	if len(strarr) 

}