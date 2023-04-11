package main

import (
	"fmt"
)

type Carro struct {
	nombre     string
	numPuertas int
}

func main() {

	carro1 := Carro{
		nombre:     "Mazda 1",
		numPuertas: 2,
	}

	carro2 := Carro{
		nombre:     "Mazda 2",
		numPuertas: 2,
	}

	carro3 := Carro{
		nombre:     "Mazda 3",
		numPuertas: 4,
	}

	carro4 := Carro{
		nombre:     "Mazda 3",
		numPuertas: 6,
	}

	carro5 := Carro{
		nombre:     "Mazda 3",
		numPuertas: 6,
	}

	arreglo := [3]Carro{
		carro1,
		carro2,
		carro3,
	}

	carrosArr := make([]Carro, 0)
	carrosArr = append(carrosArr, carro4, carro5)

	notas := map[string]float64{
		"Mario":   4.5,
		"Gerardo": 8.5,
		"Rodrigo": 3.5,
	}

	for name, calf := range notas {
		fmt.Println(name, calf)
	}

	for i := 0; i < len(carrosArr); i++ {
		fmt.Println("Nombre: ", carrosArr[i].nombre, " - Numero de Puertas: ", carrosArr[i].numPuertas)
	}

	for i := 0; i < len(arreglo); i++ {
		fmt.Println("Nombre: ", arreglo[i].nombre, " - Numero de Puertas: ", arreglo[i].numPuertas)
	}
	var miArray = [3]int{1, 2, 3}
	fmt.Println(miArray[2])
}
