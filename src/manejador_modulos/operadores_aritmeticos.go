package main

import (
	"fmt"
	"math"
)

func main() {

	a := 10
	b := 3
	//operadores aritmeticos
	fmt.Println(a / b)
	fmt.Println(a % b)
	fmt.Println(a * b)
	fmt.Println(a + b)
	fmt.Println(a - b)
	//operadores de incremento - decremento
	b++
	a--
	b = b + 2
	//suma en asignacion
	a += 1
	fmt.Println(a)
	fmt.Println(b)

	//uso del paquete math para las diferentes funciones
	fmt.Println(math.Pi)
	fmt.Println(math.E)
	fmt.Println(math.Pow(5, 2))
	fmt.Println(math.Sqrt(25))

	//raiz cubica
	fmt.Println(math.Cbrt(25))
}
