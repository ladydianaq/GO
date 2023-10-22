// es una estructura de datos
// que permite definir tipo de dato personalizado
// compuesto por diferentes campos, tipos especificos
// es un modelo de la clase
package main

import (
	"fmt"
)

// crear una estructura
type Persona struct {
	nombre string
	edad   int
	correo string
}

func main() {
	var p Persona //nstancia de la estructura personafunc
	p.nombre = "Lady"
	p.edad = 28
	p.correo = "lady@gmail.com"
	fmt.Println(p)

	//instacia en una sola linea
	//p:=Persona{"JHoselin",23,"jhoselin@gmail.com"}
	//modificar datos
	//p.edad=30

	//p2:=Persona{"JHoselin",23,"jhoselin@gmail.com"}
}
