package main

import (
	"fmt"
	"math"
)

type ErrorPersonalizado struct {
	mensaje string
}

func main() {

	// Llamada a la función multiplicar
	resultado1 := multiplicar(47, 0)
	fmt.Println("El resultado de la multiplicación es:", resultado1)

	// Llamada a la función potencia
	resultado2, err := potencia(0, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("El resultado de la potencia es:", resultado2)
	}

}

func multiplicar(a, b float64) float64 {
	if b == 0 || a == 0 {
		panic("Cualquier número multiplicado por cero da como producto cero. Esto se conoce como la propiedad cero de la multiplicación.")
	}
	return a * b
}

func (e *ErrorPersonalizado) Error() string {
	return e.mensaje
}

func potencia(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &ErrorPersonalizado{"Si la base y el exponente son 0 s considera como indeterminado"}
	}
	return math.Pow(a, b), nil
}
