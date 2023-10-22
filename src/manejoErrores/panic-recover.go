// para el manejo de errores inesperadas
package main

import "fmt"

// captura el panico dentro de la funcion
// controlar con recover
func divide(dividendo, divisor int) {
	//funciones anonimas
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)

		}

	}()
	validateZero(divisor)
	fmt.Println(dividendo / divisor)
}
func validateZero(divisor int) {
	if divisor == 0 {
		panic("no puedes dividir por cero")
	}
}
func main() {
	//panic: interrumpe si ocurre un error
	divide(100, 10)
	divide(34, 0)
}

//recover controla en panico
