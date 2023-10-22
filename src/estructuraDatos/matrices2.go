package main

import "fmt"

func main() {
	//definimos una matriz con valores predeterminados
	var matriz [5]int
	matriz[0] = 10
	matriz[1] = 20
	fmt.Println(matriz)

	//inicializar una matriz
	var matriz1 = [2]int{20, 30}
	fmt.Println(matriz1)

	//reccorrer los valores de la matriz con for
	for i := 0; i < len(matriz); i++ {
		fmt.Println(matriz[i])
	}

	//recorrer la matriz por pares utilizando range
	for index, value := range matriz {
		fmt.Printf("indice %d valor %d\n", index, value)
	}

	//definir una matriz bidimensional
	var matriz2 = [2][3]int{{1, 2, 3}, {5, 6, 9}}
	fmt.Println(matriz2)

}
