// generar un numero aleatorio entre 0 y 100
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//para que genere valores distintos tambien puedes utilizar
	//rand.Seed(time.Now().UnixNano())
	//permite obtener un numero alatorio entre 1 100
	//fmt.Println(rand.Intn(100))
	jugar()
}
func jugar() {
	nALeatorio := rand.Intn(100)
	var nIngresado int
	var intentos int
	const maxIntentos = 4

	for intentos < maxIntentos {
		intentos++
		fmt.Printf("ingrese un numero (intentos restantes: %d): ", maxIntentos-intentos+1)
		fmt.Scanln(&nIngresado)

		if nIngresado == nALeatorio {
			fmt.Println("Felicitaciones, adivinaste el numero")
			jugarNuevamente()
			return
		} else if nIngresado < nALeatorio {
			fmt.Println("El numero a adivinar es mayor")
		} else if nIngresado > nALeatorio {
			fmt.Println("El numero a adivinar es menor.")
		}

	}
	fmt.Println("se acabaron los intentos. El numero era:", nALeatorio)
	jugarNuevamente()
}

func jugarNuevamente() {
	var eleccion string
	fmt.Println("Quieres jugar nuevamente? (s/n): ")
	fmt.Scanln(&eleccion)
	switch eleccion {
	case "s":
		jugar()
	case "n":
		fmt.Println("Gracias por jugar!")
	default:
		fmt.Println("eleccion invalidad intente nuevamente.")
		jugarNuevamente()

	}

}
