package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "123f"
	num, err := strconv.Atoi(str)
	//si devuelve nil, es que no ha habido un error
	//si el err!=nil es por que ha ocurrid un error
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println("Numero: ", num)

}
