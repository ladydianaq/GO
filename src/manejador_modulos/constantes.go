package main

import (
	"fmt"
)

// declaracion e inicializacion de constantes, la declaracion de contantes son con mayusculas
const Pi float32 = 3.14

// sin utilizar el tipo de datos
const Pi2 = 3.1415

// constantes numericas
const (
	x = 100    //numerico
	y = 0b1010 //binario
	z = 0o12   //octal
	w = 0xFF   //hexadecimal
)
const (
	domingo = iota + 1
	lunes
	martes
	jueves
	viernes
	sabado
)

func main() {
	fmt.Println(x, y, z, w)

}
