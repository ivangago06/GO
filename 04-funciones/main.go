package main

import (
	"fmt"
)

func main() {
	resultado := area_rectangulo(3.5, 9.2)
	fmt.Println(resultado)
}

func suma(a int, b int) int {
	return a + b
}

func area_rectangulo(b float32, a float32) float32 {
	return (b * a) / 2
}
