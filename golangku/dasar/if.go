package main

import "fmt"

func main() {
	var angka int = 5
	if angka == 3 {
		fmt.Println("Fish")
	} else if angka == 5 {
		fmt.Println("Bush")
	} else if angka == 15 {
		fmt.Println("Fish Bush")
	}

	switch angka {
	case 1:
		fmt.Println("Ini angka 1")
	case 2, 3, 4, 5:
		fmt.Println("Ini angka 2 - 5")
	default:
		{
			fmt.Println("Ga ada yang memenuhi")
			fmt.Println("Well Played")
		}
	}

	switch {
	case angka == 1:
		fmt.Println("Well")
		fallthrough
	case angka <= 5:
		fmt.Println("Played")
	default:
		{
			fmt.Println("Well")
			fmt.Println("Played")
		}
	}

}
