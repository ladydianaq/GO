// los mapas nos pemiten trabajar con clave valor
package main

import "fmt"

func main() {
	colors := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}
	fmt.Println(colors)
	//recuperar valor
	fmt.Println(colors["rojo"])
	//agregar un nuevo elemento
	colors["negro"] = "000000"

	//imprimir una verificacion de la existencia de un elemento
	//ok la realiza la verificacion
	valor, ok := colors["blanco"]
	fmt.Println(valor, ok)

	if ok {
		fmt.Println("existe")
	} else {
		fmt.Println("no existe")
	}

	//eliminar un elementos
	delete(colors, "azul") //eliminar el elemento azul
	fmt.Println(colors)

	//iterar clave valor
	for clave, valor := range colors {
		fmt.Println("clave: %s, Valor: %s", clave, valor)
	}
}
