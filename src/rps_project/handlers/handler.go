package handlers

import (
	"fmt"
	"net/http"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Pagina de inicio")
}
