package main

import "fmt"

func main() {
	colors := map[string]string{
		"ave":       "paloma",
		"hervivoro": "oveja",
	}
	fmt.Println(colors["ave"])
	colors["negro"] = "#000000"
	valor, ok := colors["ave"]
	fmt.Println(valor)

	if ok {
		fmt.Println("existe")
	} else {
		fmt.Println("No existe")
	}

	//eliminar un elemento
	delete(colors, "negro")
	fmt.Println(colors)

	//iterar los valores con range
	for clave, valor := range colors {
		fmt.Printf("clave: %s,valor: %s \n", clave, valor)
	}
}
