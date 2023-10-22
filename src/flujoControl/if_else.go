package main

import (
	"fmt"
	"time"
)

func main() {
	// t obtiene la hora actual
	t := time.Now()
	//obtiene la hora de t
	hora := t.Hour()

	if hora < 12 {
		fmt.Println("¡Mañana!")
	} else if hora < 17 {
		fmt.Println("Tarde!")
	} else {
		fmt.Println("NOche")
	}
}
