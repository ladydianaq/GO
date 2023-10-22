package main

//definimos una tarea
type task struct {
	//se define como json, por que es como respondera al cliente
	Id      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

//allTask: creamos un tipo llamado allTask
//allTask: sera un arreglo
//cons esta expresion definimos que estamos haciendolo
//para una tarea o para una lista de tareas
type allTasks []task

var tasks =allTasks{
	Id:1,
	Name:"Lady,
	Content:"realizar un Crud"
	"

}
