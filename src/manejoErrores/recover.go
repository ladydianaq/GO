func divide(dividend, divisor int) {
	/* Se utiliza defer para posponer la ejecución de una función anónima hasta que la función que la contiene haya finalizado.
	La función anónima tiene como objetivo capturar cualquier panic que ocurra en el bloque de código que la rodea y realizar una acción de recuperación.
	Si se produce un panic, el valor recuperado se almacena en 'r' y se verifica si no es nulo.
	Si 'r' no es nulo, significa que se produjo un panic y se imprime su valor utilizando fmt.Println().
	Esta construcción es útil para manejar errores y realizar tareas de limpieza antes de que la función finalice su ejecución. */
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r) // Imprimir el valor recuperado en caso de un panic
		}
	}()

	validateZero(divisor)
	fmt.Println(dividend / divisor)