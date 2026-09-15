# Codigo GO
---

## Contrucción Basica

```Go
// Servidor Basíco
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Prueba contacto")
	})

	fmt.Println("Server run in http://localhost:800")
	http.ListenAndServe(":8080", nil)
}

// Ceparación de funciones
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Prueba contacto")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server run in http://localhost:800")
	http.ListenAndServe(":8080", nil)
}
```

## Manejo de peticiones sin mux.

```Go
func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)

	log.SetPrefix("[ventasFlow] ")
	log.Println("servidor iniciado")
	log.Printf("puerto: %d", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Pagina de Inicio")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sobre nosotros")
	log.Printf("%s %s", r.Method, r.URL.Path)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Pagina de Contacto")
	log.Printf("%s %s", r.Method, r.URL.Path)
}
```

## Manejo de rutas y parámetros sin mux.

```Go
func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/products/detail/", productDetailHandler)

	log.SetPrefix("[ventasFlow] ")
	log.Println("servidor iniciado")
	log.Printf("puerto: %d", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Pagina de Inicio")
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Información de productos")
	log.Printf("%s %s", r.Method, r.URL.Path)
}

// aqui se agrega el manejador de parametros provenientes de la URL
func productDetailHandler(w http.ResponseWriter, r *http.Request) {
	const prefix = "/products/detail/"
	if !strings.HasPrefix(r.URL.Path, prefix) {
		http.NotFound(w, r)
		return
	}
	productID := strings.TrimPrefix(r.URL.Path, prefix)
	if productID == "" {
		http.Error(w, "ID de producto requerido", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Detalles del producto con ID: %s", productID)
} 
```