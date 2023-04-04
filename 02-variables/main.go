package main

import "fmt"

var status bool = true

func main() {
	var x float32 = 3.14159265358979323846
	var y float64 = 3.14159265358979323846

	fmt.Println(x)
	fmt.Println(y)

	sum := x + float32(y)
	diff := x - float32(y)

	fmt.Println("x + y =", sum)
	fmt.Println("x + y =", diff)

	const number int = 23 // Variable no mutables

	fmt.Println(number)
	fmt.Println(status)

}
