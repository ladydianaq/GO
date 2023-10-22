//el manejo de errores permiten manejar los posibles errores que
//que pueden ocurrir durante la ejeucion

package main

import "fmt"

// el defer permite agregar las instrucciones a una pila
// el ultimo en ingresar a la pila es el primero en ejecutarse
func main() {
	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)

}
