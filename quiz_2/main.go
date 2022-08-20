package main

import "fmt"

func main() {
	for a := 1; a <= 5; a++ {
		for b := 1; b <= 6; b++ {
			if a == 1 || a == 5 || ((a == 2 || a == 5-1) && (b == 1 || b == 6)) || (a == 3 && (b == 3 || b == 4)) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
