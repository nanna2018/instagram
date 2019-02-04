package main

import (
	"fmt"
	"log"
	"net/http"
	hnd "proyecto1/handlers"
	"strconv"
)

func main() {
	port := 8080

	fmt.Println("Iniciando servidor...")

	//Primera opción
	for path, handler := range hnd.Manejadores {
		http.HandleFunc(path, handler)
	}

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
	fmt.Println("Servidor abierto en http://localhost:" + strconv.Itoa(port))
}
