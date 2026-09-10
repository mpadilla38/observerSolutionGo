package main

import (
	"fmt"
	"net/http"
)

//Contrucción Basica

/* func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Prueba contacto")
	})

	fmt.Println("Server run in http://localhost:800")
	http.ListenAndServe(":8080", nil)

} */

//Ceparación de funciones.

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Prueba contacto")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server run in http://localhost:800")
	http.ListenAndServe(":8080", nil)

}
