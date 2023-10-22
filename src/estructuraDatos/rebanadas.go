// es un tipo de datos que representa una seccion contigua
// de un arreglo
// es flexible para trabajar con los arreglos
// pemite emilinar elementos de forma dinamica
package main

import "fmt"

func main() {
	//var NombreRebanada []int
	diasSem := []string{"domingo", "LUnes", "Martes", "miercoles", "viernes", "sabado"}
	//crear otra rebada a partir de una rebanada
	diaRebanada := diasSem[0:5]
	fmt.Println(diaRebanada)
	diaRebanada = append(diaRebanada, "viernes", "otro dia")

	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada)) //los slice son dinamicos y se ven aqui

	//eliminar rebanadas
	//se elimina el indice 2
	diaRebanada = append(diaRebanada[:2], diaRebanada[3:]...)

	//crear rebanadas con make
	nombres := make([]string, 5, 10) //5 numero de elementso 10 capacidad
	fmt.Println(nombres)

	//copiar elementos de una rebanada a otra
	reb3 := []int{1, 2, 3, 5, 6}
	reb5 := make([]int, 5)
	//copiar lo elemntos de la rebanada5 a reb3
	copy(reb5, reb3)
	fmt.Println(reb5)
}
