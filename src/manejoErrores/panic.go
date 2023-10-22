package main

import "fmt"

func main() {
	divide(100, 10)
	divide(200, 25)
	divide(34, 0)
	divide(124, 8)
}

func divide(dividend, divisor int) {
	validateZero(divisor)
	fmt.Println(dividend / divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("🤕 no puedes dividir por cero")
	}
}
