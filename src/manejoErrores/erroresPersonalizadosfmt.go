package main

import "fmt"

func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("no se puede dividir por cero")
	}
	return dividendo / divisor, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Resultado:", result)
}
