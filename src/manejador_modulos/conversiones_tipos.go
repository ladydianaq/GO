package main

import (
	"fmt"
	"strconv"
)

// conversiones de datos
func main() {
	var integer16 int16 = 50  //tipo int16
	var integer32 int32 = 100 //tipo int32
	//convertimos el valor
	fmt.Println(int32(integer16) + integer32)

	s := "100"
	//Atoi convierte de un string a un entero
	i, _ := strconv.Atoi(s) //_hace el manejo de errores
	fmt.Println(i + 1)

	n := 42
	s = strconv.Itoa(n)
	fmt.Println(s + s) //concatena y da como resultado 4242

}
