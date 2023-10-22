// para el registro de mensajes
package main

import (
	"log"
	"os"
)

func main() {
	//registrar con prefijo
	log.SetPrefix("main():")
	//registrar todos los errores en un archivo
	//crear el archivo si es que no existe
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//cerrar el flujo +
	defer file.Close()
	log.SetOutput(file)
	log.Print("oye soy log")

	log.Fatal("so registro de errores") //se detine la ejecucion de nuestro programa
	log.Print("mensaje de registro")
	log.Panic(("Puedes verme"))
	log.Println(("regitro de mensaje"))
	log.SetPrefix("main():") //agrega un refijo a los mensajes de los regisros del programa

}
