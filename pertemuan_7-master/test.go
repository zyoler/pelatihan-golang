package main

import "log"

func main() {
	var Calc PerhitunganBangunDatar
	Calc = SisiPersegiPanjang{2}
	log.Println("Luas Bangun Data = ", Calc.luas())
}

type PerhitunganBangunDatar interface {
	luas() int
	keliling() int
}
type SisiPersegiPanjang struct {
	sisi int
}

func (s SisiPersegiPanjang) luas() int {
	return s.sisi * s.sisi
}

func (s SisiPersegiPanjang) keliling() int {
	return s.sisi * 4
}
