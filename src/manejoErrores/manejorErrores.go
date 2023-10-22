// mejorar la rubustez de la aplicacion
// manejo de errores en las funciones
package main

import (
	"fmt"
	"strconv"
)

func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		//return 0, errors.New("No se puedo enviar cero")
		//otra forma de control de errorres
		return 0, fmt.Errorf("No se puedo enviar cero")
	}
	return dividendo / divisor, nil
}
func main() {
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil { //si es que no ha ocurrido error
		fmt.Println("error:", err)
		return

	}
	fmt.Println("NUmero: ", num)
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	//imprimimo result si es que no existe ningun error
	fmt.Println(result)
}
