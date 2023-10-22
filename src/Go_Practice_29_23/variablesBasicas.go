/*
booleano

cadena

int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias para uint8

runa // alias para int32
     // representa un punto de código Unicode

flotador32 flotador64

complejo64 complejo128
*/

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// %T: recupera el tipo
	//$v: recupera el tipo en general
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
