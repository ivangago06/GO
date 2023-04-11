package main

import (
	"fmt"
	"math"
)

// Hacer un for infinito for{} y adentro un switch
func main() {

	var variable float32
	var variableS float32
	var variableR float64

	fmt.Scanln(&variable)
	fmt.Scanln(&variableS)
	fmt.Scanln(&variableR)

	resultado := area_rectangulo(variable, variableS)
	fmt.Println(resultado)

	resultadoS := area_circulo(variableR)
	fmt.Println(resultadoS)

}

func area_rectangulo(a float32, b float32) float32 {
	return (a * b) / 2
}

func area_circulo(r float64) float64 {
	pi := math.Pi
	cuadrado := math.Pow(r, 2)

	return pi * cuadrado
}
