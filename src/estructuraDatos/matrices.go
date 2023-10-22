// almacenan valores y tipos del mismo tipo
package main

import (
	"fmt"
)

func main() {
	//declaracion
	var m [5]int
	//modicar elementos
	m[0] = 4
	m[1] = 9
	//inicializar matriz con 5 elementos
	var a = [5]int{10, 20, 5, 3, 60}
	a[0] = 23
	a[1] = 56

	//inicializar matriz con un valor desconocido
	var a1 = [...]int{10, 20, 50, 60, 30, 40}
	fmt.Println(m)
	fmt.Println(a)
	fmt.Println(a1[1])

	//recorrer los elementos de una matriz
	for i := 0; i < len(a1); i++ {
		fmt.Println((a1[i]))

	}
	//otra forma de iterar una matriz con range
	for index, value := range a1 {
		fmt.Printf("indice: %d valor: %d\n", index, value)
	}

	//definir una matriz multidimensional
	var m3 = [3][3]int{{1, 2, 3}, {5, 4, 6}, {8, 9, 6}}
	fmt.Println(m3)

}
