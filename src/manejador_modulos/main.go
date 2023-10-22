package main

//importacion de paquetes
import (
	"fmt"
	"rsc.io/quote"
)

// Nombre del manejador de modulos: manage_moduls
func main() {
	// Obtén un mensaje de saludo e imprímelo.
	fmt.Println("Hello word")
	fmt.Println(quote.Hello())
}
