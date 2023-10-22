package main

import (
	"fmt"
	"math/rand"
)

func main() {
	jugar()
}

func jugar() {
	const maxIntentos = 2
	var valori int
	x := rand.Intn(100)
	var i int
	for i = 0; i < maxIntentos; i++ {
		fmt.Println("Ingrese un valor: ")
		fmt.Scanln(&valori)
		if valori < x {
			fmt.Println("Ingrese un valor mayor")

		} else if valori > x {
			fmt.Println("Ingrese un valor menor")

		} else if valori == x {
			fmt.Println("Felicitaciones!!!, Adivinaste el valor")
			NumIteraciones()
		}

	}
	fmt.Printf("Acabo el numero de Intentos, el valor era %d: ", x)
	NumIteraciones()
}

func NumIteraciones() {
	var intente string
	fmt.Println("Desea volver a jugar? (s/n): ")
	fmt.Scanln(&intente)
	switch intente {
	case "s":
		jugar()
	case "n":
		fmt.Println("Gracias por jugar")
	default:
		fmt.Println("valor invalido")
		NumIteraciones()

	}
}
