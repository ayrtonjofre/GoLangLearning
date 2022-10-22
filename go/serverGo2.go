package main

import (
	"net/http" //paquete para levantar server (inclyue herramientas y metodos)
)

func main() {

	// Routes
	http.HandleFunc("/", homeHandler) // Manejador de rutas en URL
	http.HandleFunc("/contact", contactHandler)

	// Start the server
	http.ListenAndServe(":8080", nil) // Habilito el puerto 8080 para la web en localhost
}

// w = write (Response para responder el request)
// r = request (metodo utilizado para requisiciones)
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Jordan Putinha"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Contact Page"))
}
