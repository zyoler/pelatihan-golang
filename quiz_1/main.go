package main

import "fmt"

func main() {
	tp := 1
	tp2 := 3
	for a := 1; a <= 7; a++ {
		for b := 1; b <= a; b++ {
			if b == a-tp && a <= tp2 {
				fmt.Print("&")
			} else if a == 4 {
				tp++
				tp2 += a
				break
			} else {
				fmt.Print("*")
			}
		}
		if a == 4 {
			continue
		}
		fmt.Println()
	}
}
