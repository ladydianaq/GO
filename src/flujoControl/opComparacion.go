package main

import (
	"fmt"
)

func main() {
	// Comparación de números
	fmt.Println(1 == 1) // true
	fmt.Println(1 != 2) // true
	fmt.Println(2 < 3)  // true
	fmt.Println(3 > 4)  // false
	fmt.Println(4 <= 4) // true
	fmt.Println(5 >= 6) // false

	// Comparación de cadenas
	fmt.Println("hola" == "hola")  // true
	fmt.Println("hola" != "adios") // true
	fmt.Println("abc" < "def")     // true
	fmt.Println("ghi" > "fgh")     // true
	fmt.Println("hij" <= "hij")    // true
	fmt.Println("klm" >= "klmno")  // false

	//comparacion de booleanos
	fmt.Println(true == true)           // true
	fmt.Println(false != true)          // true
	fmt.Println(true && false == false) // true
	fmt.Println(true || false == true)  // true

	//operadores logicos
	//AND
	x := true
	y := false
	z := x && y
	fmt.Println(z) // Imprime false

	//OR
	x1 := true
	y1 := false
	z1 := x1 || y1
	fmt.Println(z1) // Imprime true

	//NOT Logico
	x2 := true
	z2 := !x2
	fmt.Println(z2) // Imprime false

	x3 := true
	y3 := false

	// Negación
	fmt.Println(!x3) // false
	fmt.Println(!y3) // true

	// AND lógico
	fmt.Println(x3 && x3) // true
	fmt.Println(x3 && y3) // false
	fmt.Println(y3 && y3) // false

	// OR lógico
	fmt.Println(x || x) // true
	fmt.Println(x || y) // true
	fmt.Println(y || y) // false

	//Expresiones
	x4 := 5
	y4 := 10
	z4 := 15

	// Expresión con paréntesis, operadores aritméticos, de comparación y lógicos
	resultado := ((x4+y4)*z4)/(x4*y4) > z4 && x4 != y4

	// Imprimir el resultado
	fmt.Println(resultado) //False

}
