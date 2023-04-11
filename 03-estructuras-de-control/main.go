package main

import (
	"fmt"
)

type Modelo struct {
	nombre string
}

type Carro struct {
	color      string
	numPuertas int
	modelo     Modelo
}

func main() {

	carro1 := Carro{
		color:      "amarillo",
		numPuertas: 4,
		modelo: Modelo{
			nombre: "Clasico",
		},
	}

	fmt.Println(carro1.color)

	// Estructura IF
	//var status bool = true
	var number int = 10

	if number > 6 {
		fmt.Println("Alumno aprobado")
	} else {
		fmt.Println("Alumno reprobado")
	}

	// ESTRUCTURA FOR
	for i := 0; i < 10; i++ {
		//fmt.Println(i)
	}

	nombres := []string{"Rodrigo", "Gerardo", "Mario"}

	for indice, nombre := range nombres {
		fmt.Println(indice)
		fmt.Println(nombre)
	}

	var day string = "Enero"

	switch day {
	case "Lunes":
		fmt.Println("Hoy es lunes")
	case "Martes":
		fmt.Println("Hoy es Martes")
	default:
		fmt.Println("Día no valido")
	}
}
