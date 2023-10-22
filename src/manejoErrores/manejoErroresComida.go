package main

import (
	"errors"
	"fmt"
	"strconv"
)

// errNotFound representa el error de elemento no encontrado.
var errNotFound = errors.New("no encontrado")

// food es un mapa que contiene alimentos representados por emojis.
var food = map[int]string{
	1: "🍕",
	2: "🍔",
}

func main() {
	// Realizar una búsqueda con la clave "34".
	encontrado, err := buscar("34")
	if errors.Is(err, errNotFound) {
		fmt.Println("Se pudo manejar el error correctamente")
		return
	}
	if err != nil {
		fmt.Println("buscar()", err)
		return
	}

	fmt.Println(encontrado)
}

// buscar realiza una búsqueda de un emoji de comida utilizando una clave.
func buscar(clave string) (string, error) {
	// Convertir la clave a un número entero.
	num, err := strconv.Atoi(clave)
	if err != nil {
		return "", fmt.Errorf("strconv.Atoi(): %w", err)
	}

	// Encontrar el emoji de comida correspondiente al número.
	emoji, err := encontrarComida(num)
	if err != nil {
		return "", fmt.Errorf("encontrarComida(): %w", err)
	}

	return emoji, nil
}

// encontrarComida busca el emoji de comida correspondiente a un identificador.
func encontrarComida(id int) (string, error) {
	// Verificar si el identificador existe en el mapa de alimentos.
	valor, existe := food[id]
	if !existe {
		return "", errNotFound
	}

	return valor, nil
}
