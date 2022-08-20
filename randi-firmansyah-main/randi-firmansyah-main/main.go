package main

import (
	"log"
	"net/http"
	product "randi_firmansyah/controllers/productController"
	user "randi_firmansyah/controllers/userController"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Get("/product", product.GetSemuaProduct)
		r.Get("/product/{id}", product.GetProductById)
		r.Post("/product", product.PostProduct)
		r.Put("/product/{id}", product.UpdateProduct)
		r.Delete("/product/{id}", product.DeleteProduct)
	})

	r.Group(func(r chi.Router) {
		r.Get("/user", user.GetSemuaUser)
		r.Get("/user/{id}", user.GetUserById)
		r.Post("/user", user.PostUser)
		r.Put("/user/{id}", user.UpdateUser)
		r.Delete("/user/{id}", user.DeleteUser)
	})

	log.Println("Running Service")

	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Println("Error Starting Service")
	}
}
