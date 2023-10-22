// tipos de datos basicos
package main

import (
	"fmt"
	"math"
)

func main() {
	//tipos de datos numericos
	var integer int
	var integer2 uint //solo almacena valores positivos
	fmt.Println(integer)
	fmt.Println((integer2))

	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxUint16)

	//tipos de datos booleanos
	var value bool = true
	var valur2 bool = false
	fmt.Println(value)
	fmt.Println(valur2)

	//tipos de datos string
	fullName := "Lady Diana \t(alias \"roelcode\")"
	fmt.Println(fullName)

	//byte
	var a byte = 'a'
	fmt.Println(a)

	//tipo de dato rune
	//rune nos permite trabajar con valores unicode
	var r rune = '☻'
	fmt.Println(r) //nos da como resutltado un valor byte

}
