/*
Desarrolla una aplicación que permita a los usuarios guardar y administrar contactos en un archivo. Asegúrate de manejar los errores que puedan ocurrir al agregar y mostrar contactos.

Este es un ejemplo básico de un gestor de contactos que utiliza archivos para almacenar los contactos en formato JSON. El programa ofrece las opciones de agregar un nuevo contacto, mostrar todos los contactos y salir. Los contactos se guardan en el archivo "contacts.json" y se cargan al inicio del programa.

Ten en cuenta que este ejemplo solo aborda la funcionalidad básica del gestor de contactos y no incluye todas las validaciones de entrada y manejo de errores que podrías implementar. Puedes mejorar y expandir este código según tus necesidades y requisitos específicos.

Recuerda crear el archivo "contacts.json" en el mismo directorio antes de ejecutar el programa, o asegúrate de tener permisos de escritura en el directorio para crearlo automáticamente.
*/
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Estructura de contactos
type Contact struct {
	Name  string
	Email string
	Phone string
}

func main() {
	// Slice de contactos
	var contacts []Contact

	// Cargar contactos existentes desde el archivo
	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Println("Error al cargar los contactos:", err)
	}

	// Crear instancia de fubio
	reader := bufio.NewReader(os.Stdin)

	for {
		// Mostrar opciones al usuario
		fmt.Print("==== GESTOR DE CONTACTOS ====\n",
			"1. Agregar un contacto\n",
			"2. Mostrar todos los contactos\n",
			"3. Salir\n",
			"Elige una opción: ")
		// Leer la opción del usuario
		var option int
		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opción:", err)
			return
		}

		// Manejar la opción del usuario
		switch option {
		case 1:
			// Ingresar y crear contacto
			var c Contact
			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Teléfono: ")
			c.Phone, _ = reader.ReadString('\n')
			// Agregar un contacto a Slice
			contacts = append(contacts, c)

			// Guardar en un archivo json
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error al guardar el contacto:", err)
			}

		case 2:
			// Mostrar todos los contactos
			fmt.Println("=======================================")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre: %s Email: %s Telefono: %s\n",
					index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("=======================================")
		case 3:
			// Salir del programa
			return
		default:
			fmt.Println("Opción no válida")
		}

	}
}

// Guarda contatos en un archivo json
func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}
	return nil
}

// Carga contactos desde un archivo json
func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}

	return nil
}
