package main

import (
	"encoding/json"
	"log"
	"net/http"
)

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
	log.Println("Servidor corriendo en :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

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

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CrearUsuarioResponse{
		Status: http.StatusCreated,
		ID:     1,
		Nombre: req.Nombre,
	})

	log.Printf("InsertUser Status: %d", http.StatusCreated)
}
