package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("Gorutine 1: ", i)
		}
	}() // Estos paréntesis serefiere a la llamada de una función anónima.

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("Gorutine 2: ", i)
		}
	}() // Estos paréntesis serefiere a la llamada de una función anónima.

	wg.Wait()
	fmt.Println("Done")
}
