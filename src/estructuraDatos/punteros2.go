package main

import "fmt"

func main() {
	var x int = 10
	fmt.Println(x)
	hola(&x)
	fmt.Println(x)
}

func hola(x *int) {
	*x = 20
}
