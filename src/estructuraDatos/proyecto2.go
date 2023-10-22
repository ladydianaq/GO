package main

import (
	"bufio"
	"fmt"
	"os"
)

// creamos la estructura
type Tarea struct {
	titulo      string
	descripcion string
	estado      bool
}

// creamos el metodo ListaTareas
type ListaTareas struct {
	tareas []Tarea //slice tareas de tipo Tarea
}

// metodos para eliminar, agregar  y actualizar
// marcar como completado
func (e *ListaTareas) eliminarTarea(index int) {
	e.tareas = append(e.tareas[:index], e.tareas[index+1:]...)
}

func (a *ListaTareas) agregarTareas(t Tarea) {
	a.tareas = append(a.tareas, t)

}
func (m *ListaTareas) marcarCompletado(index int) {
	m.tareas[index].estado = true
}

func (ed ListaTareas) editarTareas(index int, t Tarea) {
	ed.tareas[index] = t
}

// entrada de datos por consola
func main() {
	lista := ListaTareas{}

	//Instanciar de bufio para la entrada de datos
	leer := bufio.NewReader(os.Stdin)

	//i: indice, t:tarea
	for {
		var opcion int
		fmt.Println("seleccione la opcion: \n",
			"1. Agregar tarea: \n",
			"2. Marcar como completado: \n",
			"3. Editar tarea: \n",
			"4. Eliminar tarea: \n",
			"5. Salir")

		fmt.Print("Ingrese la opcion: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.titulo, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la descripcion de la tarea: ")
			t.descripcion, _ = leer.ReadString('\n')
			lista.agregarTareas(t)
			fmt.Println("tarea agregada correctamente")
		case 2:
			var index int
			fmt.Print("Ingrese el indice de  la tarea que desea marcar como completada")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println("Tarea marcada como completada correctamente.")
		case 3:
			var index int
			var t Tarea
			fmt.Print("Ingrese el indice de la tarea que desea actualizar: ")
			fmt.Scanln(&index)
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.titulo, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la descripcion de la tarea: ")
			t.descripcion, _ = leer.ReadString('\n')
			lista.editarTareas(index, t)
			fmt.Println("Taraea actualizada correctamente")
		case 4:
			var index int
			fmt.Print("Ingresa el indice que desea eliminar: ")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Println("Tarea eliminada correctamente. ")
		case 5:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opcion invalida.")
		}

		fmt.Println("Lista de tareas")
		fmt.Println("===================================================================")
		for i, t := range lista.tareas {
			fmt.Printf("%d. %s - %s completado: %t\n ", i, t.titulo, t.descripcion, t.estado)
		}
		fmt.Println("===================================================================")

	}

}
