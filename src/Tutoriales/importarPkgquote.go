package main

import (
	"fmt"

	"rsc.io/quote"
)

// importar el paquete quote
func main() {
	//fmt.Println("Hola mundo")
	fmt.Println(quote.Go()) //nos imprime hola mundo
	fmt.Println(quote.Hello())
}
