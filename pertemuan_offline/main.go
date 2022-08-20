package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Starting Go Routine")

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go FungsiGo(&wg)
	}

	for x := 0; x < 5; x++ {
		wg.Add(1)
		go func() {
			fmt.Println("Testing 3")
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Ended go Routine")
}

func FungsiGo(wg *sync.WaitGroup) {
	fmt.Println("Testing 1")
	defer wg.Done()
}
