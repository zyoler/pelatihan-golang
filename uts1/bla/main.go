package main

import "fmt"

func main() {

	var inp, c, d, e int
	fmt.Println("Input : ")
	fmt.Scanln(&inp)
	c = inp
	d = inp

	for a := 1; a < inp*2; a++ {
		d = c
		if a < inp {
			for b := 1; b <= (inp*2)-a; b++ {
				if a > b {
					fmt.Print("-")
				} else {
					if b < inp {
						fmt.Print(d)
						d--
					} else {
						fmt.Print(d)
						d++
					}
				}
			}
			fmt.Println()
			c--
		} else {
			c++
			for b := 1; b <= inp*2-e-5; b++ {
				if a+b < inp*2 {
					fmt.Print("-")
				} else {
					if b < inp {
						fmt.Print(d)
						d--
					} else {
						fmt.Print(d)
						d++
					}
				}
			}
			fmt.Println()
			e--
		}
	}

}
