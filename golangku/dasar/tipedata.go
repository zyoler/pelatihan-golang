package main

import "fmt"

func main() {
	// var firstName string = "Dani"
	// var lastName string
	// lastName = "Hidayat"

	// var firstName = "Dani"
	// var lastName = "Hidayat"

	// firstName := "Dani"
	// lastName := "Hidayat"

	// var firstName, lastName string
	// firstName, lastName = "Dani", "Hidayat"

	// var firstName, lastName string = "Dani", "Hidayat"

	// firstName, lastName := "Dani", "Hidayat"

	firstName, lastName, age := "Dani", "Hidayat", 21

	_, prodi := "Dani", "D3 Manajemen Informatika"

	const EndPoint = "/api/"

	fmt.Println(EndPoint)

	fmt.Print("Nama saya ", firstName, " ", lastName, "\n")
	fmt.Printf("Nama saya %s %s\n", firstName, lastName)
	fmt.Println("Nama saya", firstName, lastName, "berumur", age, prodi)
}
