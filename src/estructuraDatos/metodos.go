package main

import "fmt"

// definir una estructura
type Persona struct {
	nombre string
	edad   int
	correo string
}

func (p *Persona) diHola() {
	fmt.Println("Hola, mi nombre es ", p.nombre)

}

func main() {
	p := Persona{"Lady", 26, "lady@gmail.com"}
	//para llamar al metodo de la estructura
	p.diHola()
}
