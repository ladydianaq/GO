// utilizado para posponer la ejecucion de la funcion
// hasta que la funcion que la contiene haya finalizado
package main

import (
	"fmt"
	"os"
)

func main() {
	//se realiza el control de errores en la parte de creacion de archivos
	file, err := os.Create("Hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.Write([]byte("hola, lady diana"))
	if err != nil {
		fmt.Println(err)
		return
	}
	//file.Close() //cerrar el flujo
}
