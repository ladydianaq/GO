package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		if i == 5 {
			break
		}
	}

	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue //salta a la siguiente iteracion
		}
		fmt.Println(i)
	}
}
