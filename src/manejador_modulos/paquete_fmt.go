package main

import (
	"fmt"
)

func main() {
	name := "Lady"
	age := 28

	fmt.Printf("Hola me llamo %s tengo %d \n", name, age)
	greeting := fmt.Sprintf("Hola me llamo %s tengo %d \n", name, age)
	fmt.Println(greeting)

	//escanear solo un dato
	var name1 string
	var age1 int
	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&name1)
	fmt.Println("Ingrese su edad: ")
	fmt.Scanln(&age1)

	//escanear variso datos
	var name2, name3 string
	var age2 int
	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&name2, &name3)
	fmt.Println("Ingrese su edad: ")
	fmt.Scanln(&age2) //solo te permite escanear un dato
	//v te permite imprimir un dato en el que no sabemos el tipo de dat que se va a mostrar
	fmt.Printf("Hola me llamo %v tengo %v y trabajo con %v\n", name2, name3, age2)
	fmt.Printf("tipo de dato: %T\n", name)
}
