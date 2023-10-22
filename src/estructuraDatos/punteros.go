package main

import "fmt"

func main() {
	var x int = 10
	var p *int = &x
	fmt.Println(x)
	fmt.Println(&x) //referencia de la memoria
	fmt.Println(p)
	editar(&x)
	//est es posible, por que estamos trabajando con
	//punteros
	fmt.Println(x)

}
func editar(x *int) {
	//el puntero de la memoria
	//lo que almacena x
	*x = 20
}
