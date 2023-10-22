package main

import (
	"log"
	"os"
)

func main() {
	/*
			os.O_CREATE|os.O_APPEND|os.O_WRONLY: Estas son las banderas de modo utilizadas para el archivo. En este caso, se combinan tres banderas utilizando el operador OR (|):
		os.O_CREATE: Crea el archivo si no existe.
		os.O_APPEND: Abre el archivo en modo adjunto, lo que significa que los nuevos datos se agregarán al final del archivo.
		os.O_WRONLY: Abre el archivo en modo de solo escritura.
		0644: Es el valor octal utilizado para los permisos del archivo. En este caso, 0644 significa que el archivo tendrá permisos de lectura y escritura para el propietario y solo permisos de lectura para los demás
	*/
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("¡Oye, soy un Log!")
}
