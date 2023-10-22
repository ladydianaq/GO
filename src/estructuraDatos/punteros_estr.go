package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	correo string
}

// receptor, como referencia de memoria
// el metodo funcion mediante el receptor
func (p *Persona) diHola() {
	fmt.Println("HOla mi nombre es,", p.nombre)

}
func main() {
	p := Persona{"Alex", 25, "alex@gamil.com"}
	//acceso al metodo por medio de la instancia de lapersona
	p.diHola()
}
