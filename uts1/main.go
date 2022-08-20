package main

import (
	"fmt"
	"net/http"
	"uts_1/controller"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()
	router.Get("/gethero", controller.GetHero)
	router.Post("/posthero", controller.PostHero)
	router.Delete("/deletehero/{id}", controller.DeleteHero)
	router.Put("/updatehero/{id}", controller.UpdateHero)

	router.Get("/getbook", controller.GetBook)
	router.Post("/postbook", controller.PostBook)
	router.Delete("/deletebook/{id}", controller.DeleteBook)
	router.Put("/updatebook/{id}", controller.UpdateBook)

	fmt.Println("Service On Running")

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println("Error Starting Service")
	}
	fmt.Println("Starting Service")
}
