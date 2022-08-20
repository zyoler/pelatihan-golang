package main

import (
	"fmt"
	"net/http"
	"pertemuan_2/controller"
)

func main() {
	http.HandleFunc("/getdata", controller.GetData)
	http.HandleFunc("/postdata", controller.PostData)
	fmt.Println("Running Service")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		fmt.Println("Error Starting Service")
		return
	}
	fmt.Println("Starting Services")
}
