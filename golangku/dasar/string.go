package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Dani Hidayat", "Dani"))
	fmt.Println(strings.Split("Dani Hidayat", " "))
	fmt.Println(strings.ToLower("Dani Hidayat"))
	fmt.Println(strings.ToUpper("Dani Hidayat"))
	fmt.Println(strings.Trim("    Dani Hidayat      ", " "))
	fmt.Println(strings.ReplaceAll("Dani Dani Hidayat", "Dani", "Jeri"))

	fmt.Println(strings.Compare("85", "80"))        // return 1, 0, -1
	fmt.Println(strings.Count("Dani Hidayat", "a")) // Menghitung jumlah huruf output 3
	fmt.Println(strings.EqualFold("Dani", "Dani2"))

	fmt.Println(strings.Fields("   Dani Hidayat   "))
	fmt.Println(strings.HasPrefix("Dani", "D"))    // Apakah awalannya huruf itu? return bool
	fmt.Println(strings.HasSuffix("Hidayat", "t")) // Apakah akhirannya huruf itu? return bool
}
