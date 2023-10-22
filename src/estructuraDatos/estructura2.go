package main

import "fmt"

// definir una estructura
type Persona struct {
	nombre string
	edad   int
	correo string
}

func main() {
	var p Persona
	p.nombre = "Lady"
	p.edad = 26
	p.correo = "lady@gmail.com"
	fmt.Println(p)

	p2 := Persona{"Mayli", 26, "lady@gmail.com"}
	fmt.Println(p2)
}
