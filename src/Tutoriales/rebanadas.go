package main

import "fmt"

func main() {
	//var a[]int
	diasSemana := []string{"domingo", "Lunes", "Martes", "Miercles", "Jueves", "Viernes", "Sabado"} //slice

	//crear otra rebanada, aqui devuelve solo los 5 primero e
	//elementos del arreglo
	diaRebanada := diasSemana[0:5]
	fmt.Println("dias rebanada", diaRebanada)
	//agrega elementos a la rebanada
	diaRebanada = append(diaRebanada, "viernes", "sabado", "otro diass")
	fmt.Println("dias rebanada", diaRebanada)

	//logitud y capacidad de una rebanada
	fmt.Println(len(diaRebanada))
	//capacidad de la rebanad
	fmt.Println(cap(diaRebanada))

	//eliminar rebanada
	//solo se crea la rebada hasta el 1 sin involucrar el numero 2
	diaRebanada = append(diaRebanada[:2], diaRebanada[3:]...)
	fmt.Println("dias rebanada", diaRebanada)

	//funcion make(permite crear rebanadas)
	nombres := make([]string, 5, 10) //5 longitud, 10 capacidad
	//agregar elementos
	nombres[0] = "Alex"
	fmt.Println(nombres)

	//copiar rebanadas
	rebanada11 := []int{1, 2, 3, 6}
	rebanada1 := make([]int, 5)
	//copia los elemetos de la rebanada 1 a la rebanada11
	copy(rebanada1, rebanada11)
	fmt.Println(rebanada11)
	fmt.Println(rebanada1)

}
