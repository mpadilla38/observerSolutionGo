package main

import (
	"encoding/json"
	"log"
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

/* func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Prueba contacto")
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server run in http://localhost:800")
	http.ListenAndServe(":8080", nil)

}
*/

//Manejo de peticiones

/* func main() {
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
} */

//Manejo de rutas y parámetros

/* func main() {
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
} */

// mux básico
func main() {
	mux := http.NewServeMux()

	//Listen routs
	mux.HandleFunc("GET /{$}", homeHandler)
	mux.HandleFunc("GET /usuarios", usuariosHandler)
	mux.HandleFunc("GET /updateUsuario/{id}", updateUsuarioHandler)
	mux.HandleFunc("POST /insertUsuario", insertUsuarioHandler)

	//logging
	log.SetPrefix("[ventasFlow] ")
	log.Println("servidor iniciado")
	log.Printf("puerto: %d", 8080)

	//Start server
	//handler := withNotFoundHandler(mux)
	//log.Println("Servidor corriendo en :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// Estructura de error en Root
/* type ErrorRoot struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Path    string `json:"path"`
}

// Wrapper que envuelve el mux para interceptar rutas no encontradas
func withNotFoundHandler(mux *http.ServeMux) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, pattern := mux.Handler(r)

		if pattern == "" {
			// No hubo match con ningún patrón registrado
			log.Printf("[404] Ruta no encontrada: %s %s", r.Method, r.URL.Path)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrorRoot{
				Status:  http.StatusNotFound,
				Error:   "not_found",
				Message: "El recurso solicitado no existe",
				Path:    r.URL.Path,
			})
			return
		}

		// Si hubo match, seguimos el flujo normal
		mux.ServeHTTP(w, r)
	})
}
*/

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Pagina de inicio %s", r.URL.Path)
}

func usuariosHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Listando Usuarios")
}

func updateUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	log.Printf("Detalle update: %s", id)
}

// recibiendo solicitud POST
type CrearUsuarioRequest struct {
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
	Edad   int    `json:"edad"`
}

type CrearUsuarioResponse struct {
	Status int    `json:"status"`
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
	//Path    string `json:"path"`
}

func insertUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	var req CrearUsuarioRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Error al decodificar la solicitud",
		})
		return
	}

	defer r.Body.Close()

	if req.Nombre == "" {

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "El nombre es obligatorio",
		})
		return
	}

	//Simulamos la creación de un usuario y generamos un ID ficticio
	/* resp := CrearUsuarioResponse{
		Status: http.StatusCreated,
		ID:     1,
		Nombre: req.Nombre,
	} */

	//fmt.Fprintln(w, "usuario creado")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CrearUsuarioResponse{
		Status: http.StatusCreated,
		ID:     1,
		Nombre: req.Nombre,
	})

	log.Printf("InsertUser Status: %d", http.StatusCreated)
}
