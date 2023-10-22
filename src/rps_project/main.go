package main

import (
	"log"
	"net/http"
	"project_web/handlers"
)

func main() {

	//crear un enrutador
	//router := http.NewServeMux()
	//manejador de rutas
	//router.HandleFunc("/", handlers.Index)
	//recibe un url y una funcion handler
	//la funcion registra un controlador func, el 2do parametro obtiene la peticion del cliente
	//registra un controlador para una ruta principal
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintf(writer, "hola mundo")
	//})

	//crear un servidor mediante un puerto
	//port := ":8080"
	//fmt.Printf("servidor escuchando es http://localhost%s\n", port)
	//iniciar un servidor para la apliacion y recive un puerto, manejador de rutas
	//http.ListenAndServe(port, nil)

	//router := http.NewServeMux()
	//router.HandleFunc("/", handlers.Index)
	//port := ":8080"
	//log.Printf("servidor escuchando es http://localhost%s\n", port)
	//log.Fatal(http.ListenAndServe(port, router))


	router := http.NewServeMux()
	router.HandleFunc("/", handlers.Index)
	port := ":8080"
	log.Printf("servidor escuchando es http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))

}
