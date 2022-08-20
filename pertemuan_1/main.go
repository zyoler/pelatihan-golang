package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Hello World")

	// // Variable
	// stringvar := "PUB PASIM"
	// var qty int = 123

	// fmt.Println("Isi variable", stringvar)
	// fmt.Println("Isi variable qty ", qty)

	// // Selection

	// a := 1

	// if a == 1 {
	// 	fmt.Println("Isinya 1")
	// } else {
	// 	fmt.Println("Isinya bukan 1")
	// }

	// switch a {
	// case 1:
	// 	fmt.Println("Isinya 1")
	// default:
	// 	fmt.Println("Isinya bukan 1")
	// }

	// // Looping
	// for a := 1; a <= 5; a++ {
	// 	fmt.Println("Hello World", a)
	// }

	// i := 0
	// for {
	// 	i++
	// 	fmt.Println("Inviniate loop", i)
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// // Array
	// var nama [5]string
	// nama[0] = "Dani1"
	// nama[1] = "Dani2"
	// nama[2] = "Dani3"
	// nama[3] = "Dani4"
	// nama[4] = "Dani5"
	// var namaku = [4]string{"Dani1", "Dani2", "Dani3", "Dani4"}

	// fmt.Println(nama)
	// fmt.Println(namaku)

	// // Function
	// Buah(1, 1000, "Mangga")
	// buah1()
	// fmt.Println("Calculate func", Calculate(1, 2))
	// integervar := TambahKata("Dani", "Hidayat")
	// fmt.Println("Isi integervar", integervar)
	// tamp1, tamp2 := CalculatePerkalianPertambahan(1, 2)
	// fmt.Println(tamp1, tamp2)
	// fmt.Println(CalculatePerkalianPertambahan(1, 2))
	// var tamp [2]int
	// tamp[0], tamp[1] = CalculatePerkalianPertambahan(1, 2)

	// Struct
	// Declare Struct
	type MhsPasim struct {
		nim           string
		namaMahasiswa string
	}
	var PubMhs MhsPasim // Declare Object
	// var PubMhs2 = MhsPasim{}
	// Assign value
	PubMhs.namaMahasiswa = "Dani Hidayat"
	PubMhs.nim = "02042011011"
	fmt.Println("Nama :", PubMhs.namaMahasiswa, "Nim :", PubMhs.nim)

	var AllMhs = []MhsPasim{
		{nim: "0204", namaMahasiswa: "Dani"},
		{nim: "0204", namaMahasiswa: "Dani"},
		{nim: "0204", namaMahasiswa: "Dani"},
		{nim: "0204", namaMahasiswa: "Dani"},
		{nim: "0204", namaMahasiswa: "Dani"},
	}
	for i := 0; i < len(AllMhs); i++ {
		fmt.Println("Mahasiswa ke", i+1)
		fmt.Println("Nama : ", AllMhs[i].namaMahasiswa)
		fmt.Println("Nim : ", AllMhs[i].nim)
	}

	var mhspub = MhsPasim{nim: "0204", namaMahasiswa: "Dani"}
	fmt.Println(mhspub)
}

func Buah(qty, harga int, nama string) {
	// fmt.Println("Ini function public")
	// fmt.Println("Buah", nama, "jumlahnya", qty, "Harga", harga)
	fmt.Printf("Buah besar %s jumlahnya %d dengan harga %d \n", nama, qty, harga)
}

func buah1() {
	fmt.Println("Ini function private")
	return
	fmt.Println("Test")
}

func Calculate(value1, value2 int) int {
	return value1 * value2
}

func TambahKata(value1, value2 string) string {
	return value1 + value2
}

func CalculatePerkalianPertambahan(value1, value2 int) (int, int) {
	perkalian := value1 * value2
	penjumlahan := value1 + value2
	return perkalian, penjumlahan
}
