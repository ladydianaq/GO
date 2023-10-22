// declarar variables
package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	//fmt.Println("Hola mundo")
	fmt.Println(quote.Go()) //nos imprime hola mundo
	fmt.Println(quote.Hello())

	//DECLARAR UNA VARIABLE
	var fName string //declara una variable
	var n1, n2 int   //declarar varias variabbles

	//asignar los valores
	fName = "Lady"
	n1 = 26
	n2 = 23

	fmt.Println(fName, n1, n2)

	//otra forma de declarar variables
	var (
		firstName, lastName string
		age                 int
	)
	firstName = "Fiorela"
	lastName = "Lidia"
	age = 28
	println(firstName, lastName, age)

	//declara e inicializar las variables
	var (
		firstName1 string = "Alex"
		lastname2  string = "Alejandro"
		age2       int    = 23
	)
	println(firstName1, lastname2, age2)

	//la anterior forma podemos realizarlo en una sola linea
	var fn1, ln, a = "jhoselin", "quinto", 23
	fmt.Println(fn1, ln, a)

	//para declara variables nuevas dentro de las funciones
	fn2, ln1, a1 := "jhoselin", "quinto", 23
	//para modificar
	a1 = 24
	fmt.Println(fn2, ln1, a1)
}
