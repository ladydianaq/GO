package main

import (
	"fmt"
)

func main() {
	hello()
	hello_2("Lady")
	s := hello_3("Jhoselin")
	fmt.Println(s)
	sum, mul := calc1(4, 5)
	sum1, mul1 := calc2(8, 5)
	fmt.Println(sum)
	fmt.Println(mul)

	fmt.Println(sum1)
	fmt.Println(mul1)
}

// nos ayudan a modularizar nuetras funciones
func hello() {
	fmt.Println("HOla desde la funcion")
}

func hello_2(name string) {
	fmt.Println("HOla,", name)
}

// funcion que retorna valores
func hello_3(name string) string {
	return "Hola, " + name
}

func calc(a, b int) int {
	return a + b
}

// funcion que devuelve varios valores
func calc1(a, b int) (int, int) {
	sum := a + b
	mul := a * b
	return sum, mul
}

// funcion con multiples valores
func calc2(a, b int) (sum, mul int) {
	sum = a + b
	mul = a * b
	return
}
