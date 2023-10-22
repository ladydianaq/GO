package main

import (
	"fmt"
	"log"
)

func main() {
	log.Print("Este es un mensaje de registro")
	log.Println("Este es otro mensaje de registro")

	log.Fatal("¡Oye, soy un registro de errores!")
	fmt.Print("¿Puedes verme?")

	log.Panic("¡Oye, soy un registro de errores!")
	fmt.Print("¿Puedes verme?")

	log.SetPrefix("main(): ")
	log.Print("¡Oye, soy un Log!")
	log.Fatal("¡Oye, soy un registro de errores!")
}
