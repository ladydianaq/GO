package main

import "fmt"

func main() {
	//definir una rebanada
	var a [5]int
	fmt.Println(a)

	//inicializar una rebanada
	diasSemana := []string{"lunes", "martes", "miercoles", "jueves", "viernes", "sabado", "domingo"}
	fmt.Println(diasSemana)

	//tamaño y capacidad de la rebanada
	fmt.Println(len(diasSemana))
	fmt.Println(cap(diasSemana))

	//crear una rebanada cn make
	cReb := make([]int, 5, 10)
	cReb2 := []int{5, 56, 23, 51, 97}
	fmt.Println(cReb)
	fmt.Println(cReb2)
	copy(cReb, cReb2)
	fmt.Println(cReb)

	//agragar valores a cReb2
	a_cReb2 := append(cReb2, 81, 15, 19, 89)

	fmt.Println(a_cReb2)

	//eliminar un valor
	num := append(a_cReb2[:2], a_cReb2[3:]...)
	fmt.Println(num)

}
